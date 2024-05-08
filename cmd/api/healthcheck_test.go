package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	tests := []struct {
		name         string
		config       config
		expectedCode int
	}{
		{
			name: "Available",
			config: config{
				env: "production",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Development Environment",
			config: config{
				env: "development",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &application{
				config: tt.config,
			}

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rr := httptest.NewRecorder()

			app.healthcheckHandler(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("Got %d, expected %d", rr.Code, tt.expectedCode)
			}

			resp := rr.Result()
			defer resp.Body.Close()

			fmt.Printf("response status code: %d\n", resp.StatusCode)

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("failed to read response body: %s", err)
			}

			fmt.Println(string(data))
		})
	}
}
