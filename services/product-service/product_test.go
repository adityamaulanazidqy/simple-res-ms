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

var baseURL = "http://localhost:8082"

func TestCreateProductSuccess(t *testing.T) {
	body := model.Product{
		Name:        "Burger",
		Description: "Delicious beef burger",
		Price:       15.99,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/product", baseURL), bytes.NewBuffer(bodyJSON))
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

	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestGetAllProductsSuccess(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/product", baseURL), nil)
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

func TestGetProductByIDSuccess(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/product/1", baseURL), nil)
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

func TestGetProductByIDNotFound(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/product/999", baseURL), nil)
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

func TestUpdateProductSuccess(t *testing.T) {
	body := model.Product{
		Name:        "Updated Burger",
		Description: "Updated delicious beef burger",
		Price:       17.99,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
		return
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/product/1", baseURL), bytes.NewBuffer(bodyJSON))
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

func TestUpdateProductNotFound(t *testing.T) {
	body := model.Product{
		Name:        "Non-existent Product",
		Description: "This product doesn't exist",
		Price:       99.99,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
		return
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/product/999", baseURL), bytes.NewBuffer(bodyJSON))
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

func TestDeleteProductSuccess(t *testing.T) {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/product/1", baseURL), nil)
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

func TestDeleteProductNotFound(t *testing.T) {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/product/999", baseURL), nil)
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

func TestProductServiceMethodNotAllowed(t *testing.T) {
	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/product", baseURL), nil)
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

func TestCreateProductInvalidData(t *testing.T) {
	invalidBody := `{"name": 123, "description": null, "price": "invalid"}`

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/product", baseURL),
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
