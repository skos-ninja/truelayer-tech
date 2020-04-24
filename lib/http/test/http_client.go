package test

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
)

func CreateTestHTTPClient(handler http.HandlerFunc) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // TLS returns a valid example.com cert
			},
		},
	}

	return client, s.Close
}
