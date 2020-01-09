package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

var (
	Address     = flag.String("addr", "0.0.0.0:8080", "http service address")
	count       = 0
	mu          = sync.RWMutex{}
	tempSensors = map[string]string{
		"GenuineIntel": `coretemp-isa-0000.Package\ id\ 0.temp1_input`,
		"AuthenticAMD": `nct6798-isa-0290.SMBUSMASTER\ 0.temp8_input`,
	}
)

var upgrader = websocket.Upgrader{} // use default options
var currentMessage *Message

// Message is a json message containing cpu temps
type Message struct {
	TimeStamp    time.Time `json:"timestamp"`
	CurrentTempl float64   `json:"temp"`
	Frequency    float64   `json:"frequency"`
	MinTemp      float64   `json:"min_temp"`
	MaxTemp      float64   `json:"max_temp"`
	MeanTemp     float64   `json:"mean_temp"`
	TempMA30     float64   `json:"temp_ma30"`
	MinFreq      float64   `json:"min_freq"`
	MaxFreq      float64   `json:"max_freq"`
	MeanFreq     float64   `json:"mean_freq"`
}

func newMessage(msg *Message) *Message {
	return &Message{
		TimeStamp:    msg.TimeStamp,
		CurrentTempl: msg.CurrentTempl,
		Frequency:    msg.Frequency,
		MinTemp:      msg.MinTemp,
		MaxTemp:      msg.MaxTemp,
		MeanTemp:     msg.MeanTemp,
		MinFreq:      msg.MinFreq,
		MaxFreq:      msg.MaxFreq,
		MeanFreq:     msg.MeanFreq,
		TempMA30:     msg.TempMA30,
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	mt, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", msg)

	for {
		message := getWebSocketMessage()
		message.MeanFreq = math.Round(message.MeanFreq*100) / 100
		message.MeanTemp = math.Round(message.MeanTemp*100) / 100
		t, err := json.Marshal(message)
		if err != nil {
			log.Printf("error marshalling stuff: %+v", err)
			break
		}
		err = c.WriteMessage(mt, t)
		if err != nil {
			log.Println("write:", err)
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func getWebSocketMessage() *Message {
	mu.RLock()
	defer mu.RUnlock()
	return newMessage(currentMessage)
}

func readTempFrequence(cpuName string, startTime time.Time, maCounter int64) {
	mu.Lock()
	defer mu.Unlock()
	count++
	temp := getTemp(cpuName)
	maxFrequenct := getMaxFrequence()
	message := &Message{
		TimeStamp:    time.Now(),
		Frequency:    maxFrequenct,
		CurrentTempl: temp,
		MaxTemp:      temp,
		MinTemp:      temp,
		MeanTemp:     temp,
		TempMA30:     temp,
		MaxFreq:      maxFrequenct,
		MinFreq:      maxFrequenct,
		MeanFreq:     maxFrequenct,
	}

	if currentMessage == nil {
		currentMessage = message
	}
	if temp > currentMessage.MaxTemp {
		currentMessage.MaxTemp = temp
	}

	if temp < currentMessage.MinTemp {
		currentMessage.MinTemp = temp
	}

	if maxFrequenct > currentMessage.MaxFreq {
		currentMessage.MaxFreq = maxFrequenct
	}
	if maxFrequenct < currentMessage.MinFreq {
		currentMessage.MinFreq = maxFrequenct
	}

	tempDiff := (temp - currentMessage.MeanTemp) / float64(count)
	meanTemp := currentMessage.MeanTemp + tempDiff
	currentMessage.MeanTemp = meanTemp

	maTempDiff := (temp - currentMessage.TempMA30) / float64(maCounter)
	maMeanTemp := currentMessage.TempMA30 + maTempDiff
	currentMessage.TempMA30 = maMeanTemp

	freqDiff := (maxFrequenct - currentMessage.MeanFreq) / float64(count)
	meanFreq := currentMessage.MeanFreq + freqDiff
	currentMessage.MeanFreq = meanFreq

	currentMessage.Frequency = maxFrequenct
	currentMessage.CurrentTempl = temp
}

func getTemp(cpuName string) float64 {
	cmd := exec.Command("sensors", "-j")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("error reading sensors output: %v", err)
		return 0
	}
	sensorString, ok := tempSensors[cpuName]
	if !ok {
		log.Printf("no temp sensor found for: %s", cpuName)
		return 0
	}
	value := gjson.Get(string(output), sensorString)
	temp := value.Float()
	return math.Round(temp*100) / 100
}

func getMaxFrequence() float64 {
	raw, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Printf("error reading cpuingo: %v", err)
		return 0
	}
	output := string(raw)
	lines := strings.Split(output, "\n")
	var maxSpeed float64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			a := strings.Split(line, ":")
			key := strings.TrimSpace(a[0])
			val := strings.TrimSpace(a[1])
			if key == "cpu MHz" {
				freq, err := strconv.ParseFloat(val, 64)
				if err == nil && (freq > maxSpeed) {
					maxSpeed = math.Round(freq*100) / 100
				}
			}
		}
	}
	return maxSpeed
}

func getCPUName() string {
	raw, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Printf("error reading cpuingo: %v", err)
		return ""
	}
	output := string(raw)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			a := strings.Split(line, ":")
			key := strings.TrimSpace(a[0])
			val := strings.TrimSpace(a[1])
			if key == "vendor_id" {
				return val
			}
		}
	}
	return ""
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse(string(MustAsset("assets/home.html"))))

func main() {
	flag.Parse()
	log.SetFlags(0)
	cpuName := getCPUName()
	go func() {
		startTime := time.Now()
		var maCounter int64
		maCounter = 0
		for {
			maCounter++
			readTempFrequence(cpuName, startTime, maCounter)
			time.Sleep(500 * time.Millisecond)
			currentTime := time.Now()
			t := startTime.Add(30 * time.Second)
			// if we are here that means we are past 30s
			if currentTime.After(t) {
				startTime = time.Now()
				maCounter = 0
			}

		}
	}()

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*Address, nil))
}
