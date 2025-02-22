package integrationtest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func httpPost(url string, payload []byte, headers map[string]string, t *testing.T) APIResponse {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		t.Error(err.Error())
	}

	b, _ := json.MarshalIndent(apiResponse, "", "  ")
	t.Log(string(b))

	return apiResponse
}

func httpGet(url string, headers map[string]string, t *testing.T) APIResponse {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		t.Error(err.Error())
	}

	b, _ := json.MarshalIndent(apiResponse, "", "  ")
	t.Log(string(b))

	return apiResponse
}
