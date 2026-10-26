package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestConfiguredOrigins(t *testing.T) {
	upgrader, err := configuredUpgrader("https://fast.home.lan:9090")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name, origin string
		tls, allowed bool
	}{
		{"Cockpit", "https://fast.home.lan:9090", true, true},
		{"same origin TLS", "https://fast.home.lan:9000", true, true},
		{"same origin HTTP", "http://fast.home.lan:9000", false, true},
		{"non-browser", "", true, true},
		{"wrong scheme", "http://fast.home.lan:9090", true, false},
		{"wrong port", "https://fast.home.lan:9091", true, false},
		{"wrong host", "https://evil.example:9090", true, false},
		{"host suffix", "https://fast.home.lan.evil.example:9090", true, false},
		{"null", "null", true, false},
		{"path", "https://fast.home.lan:9090/path", true, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "http://fast.home.lan:9000/echo", nil)
			if tc.tls {
				request.TLS = &tls.ConnectionState{}
			}
			if tc.origin != "" {
				request.Header.Set("Origin", tc.origin)
			}
			if got := upgrader.CheckOrigin(request); got != tc.allowed {
				t.Errorf("CheckOrigin = %v, want %v", got, tc.allowed)
			}
		})
	}
	request := httptest.NewRequest("GET", "https://fast.home.lan:9000/echo", nil)
	request.Header.Add("Origin", "https://fast.home.lan:9090")
	request.Header.Add("Origin", "https://evil.example")
	if upgrader.CheckOrigin(request) {
		t.Fatal("accepted multiple Origin headers")
	}
}

func TestOriginConfiguration(t *testing.T) {
	for _, input := range []string{"*", "null", "https://", "https://fast.home.lan/path", "https://user@fast.home.lan", "https://fast.home.lan?", "https://fast.home.lan#"} {
		if _, err := configuredUpgrader(input); err == nil {
			t.Errorf("accepted invalid origin %q", input)
		}
	}
	upgrader, err := configuredUpgrader(" https://one.example:9090, https://two.example:9090 ")
	if err != nil {
		t.Fatal(err)
	}
	for _, origin := range []string{"https://one.example:9090", "https://two.example:9090"} {
		request := httptest.NewRequest("GET", "https://sensor.example/echo", nil)
		request.Header.Set("Origin", origin)
		if !upgrader.CheckOrigin(request) {
			t.Errorf("rejected configured origin %q", origin)
		}
	}
}

func TestTLSWebSocket(t *testing.T) {
	upgrade, err := configuredUpgrader("https://fast.home.lan:9090")
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		connection, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer connection.Close()
		kind, message, err := connection.ReadMessage()
		if err == nil {
			connection.WriteMessage(kind, message)
		}
	}))
	defer server.Close()
	roots := server.Client().Transport.(*http.Transport).TLSClientConfig.RootCAs
	dialer := websocket.Dialer{TLSClientConfig: &tls.Config{RootCAs: roots}, HandshakeTimeout: time.Second}
	endpoint := "wss" + strings.TrimPrefix(server.URL, "https")
	connection, _, err := dialer.Dial(endpoint, http.Header{"Origin": {"https://fast.home.lan:9090"}})
	if err != nil {
		t.Fatal(err)
	}
	if err := connection.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	connection.SetReadDeadline(time.Now().Add(time.Second))
	_, message, err := connection.ReadMessage()
	connection.Close()
	if err != nil || string(message) != "hello" {
		t.Fatalf("TLS message = %q, %v", message, err)
	}
	rejected, response, err := dialer.Dial(endpoint, http.Header{"Origin": {"https://evil.example"}})
	if rejected != nil {
		rejected.Close()
	}
	if response != nil {
		defer response.Body.Close()
	}
	if err == nil || response == nil || response.StatusCode != http.StatusForbidden {
		t.Fatalf("unexpected upgrade result: %v, %v", response, err)
	}
}

func TestHomeWebSocketScheme(t *testing.T) {
	for _, scheme := range []string{"http", "https"} {
		request := httptest.NewRequest("GET", scheme+"://fast.home.lan:9000/", nil)
		response := httptest.NewRecorder()
		home(response, request)
		body, err := io.ReadAll(response.Result().Body)
		if err != nil {
			t.Fatal(err)
		}
		expected := "ws://fast.home.lan:9000/echo"
		if scheme == "https" {
			expected = "wss://fast.home.lan:9000/echo"
		}
		match := regexp.MustCompile(`new WebSocket\(("(?:\\.|[^"\\])*")\)`).FindSubmatch(body)
		var endpoint string
		if len(match) != 2 {
			t.Fatal("home missing WebSocket constructor")
		}
		if err := json.Unmarshal(match[1], &endpoint); err != nil {
			t.Fatal(err)
		}
		if endpoint != expected {
			t.Errorf("home WebSocket = %q, want %q", endpoint, expected)
		}
	}
}

func TestTLSRequiresBothFiles(t *testing.T) {
	for _, pair := range [][2]string{{"cert", ""}, {"", "key"}} {
		if err := serve("127.0.0.1:0", pair[0], pair[1], nil); err == nil {
			t.Fatal("missing TLS file accepted")
		}
	}
}

func TestEchoCloseHandshake(t *testing.T) {
	previousUpgrader, previousMessage := upgrader, currentMessage
	upgrader, currentMessage = websocket.Upgrader{}, &Message{CurrentTempl: 32}
	defer func() { upgrader, currentMessage = previousUpgrader, previousMessage }()
	finished := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer close(finished)
		echo(w, r)
	}))
	defer server.Close()
	connection, _, err := websocket.DefaultDialer.Dial("ws"+strings.TrimPrefix(server.URL, "http"), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer connection.Close()
	if err := connection.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	connection.SetReadDeadline(time.Now().Add(2 * time.Second))
	if _, _, err := connection.ReadMessage(); err != nil {
		t.Fatal(err)
	}
	deadline := time.Now().Add(2 * time.Second)
	if err := connection.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), deadline); err != nil {
		t.Fatal(err)
	}
	_, _, err = connection.ReadMessage()
	if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		t.Fatalf("expected close handshake, got %v", err)
	}
	select {
	case <-finished:
	case <-time.After(2 * time.Second):
		t.Fatal("sensor handler did not exit after client close")
	}
}
