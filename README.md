# OCDSensor

Collects CPU temperatures and frequencies and streams JSON samples at `/echo`.
WebSocket clients send `hello` to begin receiving samples. The default listener
is plain HTTP; provide both TLS flags to enable HTTPS and WSS.

For the Cockpit plugin served at `https://fast.home.lan:9090`, use:

```sh
ocdsensor --addr=0.0.0.0:9000 \
  --tls-cert=/etc/cockpit/ws-certs.d/0-self-signed.cert \
  --tls-key=/etc/cockpit/ws-certs.d/0-self-signed.key \
  --allowed-origins=https://fast.home.lan:9090
```

This reuses Cockpit's certificate and key in place. The service account must
already have permission to read the private key; the supplied systemd unit runs
as root. Do not make the key readable by other users. Restart OCDSensor after
the certificate files are replaced. Clients must trust the certificate or its
issuing CA.

`--allowed-origins` accepts a comma-separated list of exact HTTP(S) origins,
including scheme and port. It adds those origins to same-origin access. Other
browser origins are rejected. Non-browser clients without an Origin header
remain supported. This controls browser origins; it does not authenticate
clients. The service continues to expose read-only sensor readings.

Build and run the focused package tests:

```sh
go build -o bin/ocdsensor .
go test -v .
```

The `ocdsensor.service` template matches this machine's TLS configuration.
