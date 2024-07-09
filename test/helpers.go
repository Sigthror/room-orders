package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
)

const (
	hostEnv = "APP_HOST"
	portEnv = "APP_PORT"
)

func getHostPort() (string, error) {
	host := os.Getenv(hostEnv)
	if host == "" {
		return "", errors.New("empty host ENV variable")
	}
	port := os.Getenv(portEnv)
	if port == "" {
		return "", errors.New("empty port ENV variable")
	}

	return net.JoinHostPort(host, port), nil
}

func makeRequest(method string, url url.URL, data any) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	req, err := http.NewRequest(method, url.String(), bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}

	return resp, nil
}
