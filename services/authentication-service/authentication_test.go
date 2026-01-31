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
		Username: "nonexistentuser",
		Password: "testpassword",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/login", baseURL), bytes.NewBuffer(bodyJSON))
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

func TestRegisterAndLoginUserSuccess(t *testing.T) {
	registerBody := model.UserRegisterRequest{
		Username: "testuser_success",
		Password: "testpassword",
	}

	registerBodyJSON, err := json.Marshal(registerBody)
	if err != nil {
		t.Errorf("Error marshaling register request body: %v", err)
	}

	registerRequest, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/register", baseURL),
		bytes.NewBuffer(registerBodyJSON),
	)
	if err != nil {
		t.Errorf("Error creating register request: %v", err)
		return
	}

	client := &http.Client{}
	registerResponse, err := client.Do(registerRequest)
	if err != nil {
		t.Errorf("Error making register request: %v", err)
		return
	}
	defer registerResponse.Body.Close()

	if registerResponse.StatusCode != http.StatusOK {
		t.Errorf("Expected register status code %d, but got %d", http.StatusOK, registerResponse.StatusCode)
	}

	loginBody := model.UserLoginRequest{
		Username: "testuser_success",
		Password: "testpassword",
	}

	loginBodyJSON, err := json.Marshal(loginBody)
	if err != nil {
		t.Errorf("Error marshaling login request body: %v", err)
		return
	}

	loginRequest, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/login", baseURL),
		bytes.NewBuffer(loginBodyJSON),
	)
	if err != nil {
		t.Errorf("Error creating login request: %v", err)
		return
	}

	loginResponse, err := client.Do(loginRequest)
	if err != nil {
		t.Errorf("Error making login request: %v", err)
		return
	}
	defer loginResponse.Body.Close()

	if loginResponse.StatusCode != http.StatusOK {
		t.Errorf("Expected login status code %d, but got %d", http.StatusOK, loginResponse.StatusCode)
	}

	loginBodyBytes, err := io.ReadAll(loginResponse.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Login response status code: %d", loginResponse.StatusCode)
	t.Logf("Login response body: %s", string(loginBodyBytes))
}

func TestRegisterUserAlreadyExists(t *testing.T) {
	client := &http.Client{}

	body := model.UserRegisterRequest{
		Username: "duplicateuser",
		Password: "testpassword",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/register", baseURL), bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer response.Body.Close()

	duplicateBodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Error marshaling duplicate request body: %v", err)
	}

	duplicateRequest, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/register", baseURL),
		bytes.NewBuffer(duplicateBodyJSON),
	)
	if err != nil {
		t.Errorf("Error creating duplicate request: %v", err)
		return
	}

	duplicateResponse, err := client.Do(duplicateRequest)
	if err != nil {
		t.Errorf("Error making duplicate request: %v", err)
		return
	}
	defer duplicateResponse.Body.Close()

	if duplicateResponse.StatusCode != http.StatusConflict {
		t.Errorf("Expected status code %d, but got %d", http.StatusConflict, duplicateResponse.StatusCode)
	}

	bodyBytes, err := io.ReadAll(duplicateResponse.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", duplicateResponse.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}

func TestInvalidLoginCredentials(t *testing.T) {
	registerBody := model.UserRegisterRequest{
		Username: "credentialtest",
		Password: "correctpassword",
	}

	registerBodyJSON, err := json.Marshal(registerBody)
	if err != nil {
		t.Errorf("Error marshaling register request body: %v", err)
	}

	client := &http.Client{}

	registerRequest, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/register", baseURL),
		bytes.NewBuffer(registerBodyJSON),
	)
	if err != nil {
		t.Errorf("Error creating register request: %v", err)
		return
	}

	registerResponse, err := client.Do(registerRequest)
	if err != nil {
		t.Errorf("Error making register request: %v", err)
		return
	}
	defer registerResponse.Body.Close()

	loginBody := model.UserLoginRequest{
		Username: "credentialtest",
		Password: "wrongpassword",
	}

	loginBodyJSON, err := json.Marshal(loginBody)
	if err != nil {
		t.Errorf("Error marshaling login request body: %v", err)
		return
	}

	loginRequest, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/login", baseURL),
		bytes.NewBuffer(loginBodyJSON),
	)
	if err != nil {
		t.Errorf("Error creating login request: %v", err)
		return
	}

	loginResponse, err := client.Do(loginRequest)
	if err != nil {
		t.Errorf("Error making login request: %v", err)
		return
	}
	defer loginResponse.Body.Close()

	if loginResponse.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, but got %d", http.StatusUnauthorized, loginResponse.StatusCode)
	}

	bodyBytes, err := io.ReadAll(loginResponse.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	t.Logf("Response status code: %d", loginResponse.StatusCode)
	t.Logf("Response body: %s", string(bodyBytes))
}
