package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	testCases := []struct {
		name         string
		config       config
		expectedCode int
	}{
		{
			name: "development",
			config: config{
				env: "development",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "production",
			config: config{
				env: "production",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range testCases {
		t.Run("test development healthcheck", func(t *testing.T) {
			app := application{
				config: tt.config,
			}

			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)

			app.healthcheckHandler(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("Expected code: %d, got code: %d\n", tt.expectedCode, rr.Code)
			}

			resp := rr.Result()
			defer resp.Body.Close()

			fmt.Printf("response status code: %d\n", resp.StatusCode)

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error reading from response body: %s", err)
			}

			fmt.Println(string(data))
		})
	}
}
