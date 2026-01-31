package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restaurant/model"
	"testing"
)

var baseURL = "http://localhost:8080"

func TestCreateOrderUserNotFound(t *testing.T) {
	body := model.OrderRequest{
		UserID:     999,
		ProductID:  1,
		Quantity:   2,
		TotalPrice: 50.0,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/order", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestCreateOrderSuccess(t *testing.T) {
	body := model.OrderRequest{
		UserID:     1,
		ProductID:  2,
		Quantity:   2,
		TotalPrice: 25.0,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/order", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestGetOrdersSuccess(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/order", baseURL), nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code %d or %d, but got %d", http.StatusOK, http.StatusNoContent, response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestCreateOrderInvalidData(t *testing.T) {
	invalidBody := `{"user_id": "invalid", "product_id": "invalid", "quantity": "invalid", "total_price": "invalid"}`

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/order", baseURL),
		bytes.NewBufferString(invalidBody),
	)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestOrderServiceMethodNotAllowed(t *testing.T) {
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/order", baseURL), nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}
