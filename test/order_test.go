//go:build integration

package test

import (
	"net/http"
	"net/url"
	"testing"
	"time"
	"top-selection-test/internal/model"
)

func TestOrder(t *testing.T) {
	// host, err := getHostPort()
	// if err != nil {
	// 	t.Fatalf("failed to get host: %v", err)
	// }
	tests := []struct {
		name   string
		method string
		data   any
		code   int
	}{
		{
			name:   "unexpected json field",
			method: http.MethodPost,
			data: struct {
				Key string `json:"key"`
			}{
				Key: "value",
			},
			code: http.StatusBadRequest,
		},
		{
			name:   "success creating order",
			method: http.MethodPost,
			data: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      time.Now(),
				To:        time.Now().AddDate(0, 0, 3),
			},
			code: http.StatusOK,
		},
	}

	url := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "api/v1/order",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := makeRequest(tt.method, url, tt.data)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != tt.code {
				t.Fatalf("wrong http code: want %d have %d", tt.code, resp.StatusCode)
			}
		})
	}
}
