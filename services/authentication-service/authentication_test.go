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

var baseURL = "http://localhost:8081"

func TestLoginUserNotFound(t *testing.T) {
	body := model.UserLoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/login", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, response.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(response.Body)

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestLoginUserSuccess(t *testing.T) {
	body := model.UserLoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/login", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(response.Body)

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestRegisterUserSuccess(t *testing.T) {
	body := model.UserRegisterRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/register", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(response.Body)

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestRegisterUserAlreadyExists(t *testing.T) {
	body := model.UserRegisterRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/register", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusConflict {
		t.Errorf("Expected status code %d, but got %d", http.StatusConflict, response.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(response.Body)

	t.Logf("Response status code: %d", response.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}
