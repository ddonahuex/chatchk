package knowledge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"nethopper.io/utils"
)

// FileResponse represents the response from the file upload endpoint
type FileResponse struct {
	ID string `json:"id"`
}

// KnowledgeResponse represents the response from the knowledge base creation endpoint
type KnowledgeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Detail string `json:"detail"`
}

// Endpoint: POST /api/v1/files/
//
//	OpenAI API compatible file upload endpoint.
func uploadFile(ip, port, apiKey, filePath string) (string, error) {
	// api endpoint
	endpoint := "api/v1/files/"

	// locals
	var err error
	var errMsg string

	// get Ollama ip:port & token from env
	err = utils.GetOllamaEnvVars(&ip, &port, &apiKey)
	if err != nil {
		return "", err
	}

	// Create buffer for multipart form
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		errMsg = fmt.Sprintf("Error opening %s: %v", filePath, err)
		return "", errors.New(errMsg)
	}
	defer file.Close()

	// Create form file field
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		errMsg = fmt.Sprintf("Error writing file from %s: %v", filePath, err)
		return "", errors.New(errMsg)
	}

	// Copy file content to form
	_, err = io.Copy(part, file)
	if err != nil {
		errMsg = fmt.Sprintf("Error opying content to form: %v", err)
		return "", errors.New(errMsg)
	}

	// Close multipart writer
	writer.Close()

	// Create HTTP request
	url := fmt.Sprintf("http://%s:%s/%s", ip, port, endpoint)
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		errMsg = fmt.Sprintf("Error creating HTTP request: %v", err)
		return "", errors.New(errMsg)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg = fmt.Sprintf("HTTP client error: %v", err)
		return "", errors.New(errMsg)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return "", fmt.Errorf("failed to decode error response: %w", err)
		}
		return "", fmt.Errorf("file upload failed with status %d: %s", resp.StatusCode, errResp.Detail)
	}

	var fileResp FileResponse
	if err := json.NewDecoder(resp.Body).Decode(&fileResp); err != nil {
		return "", fmt.Errorf("failed to decode file response: %w", err)
	}

	// if here, all is well
	return fileResp.ID, nil
}

// createKnowledgeBase creates a new knowledge base in Open WebUI and returns the knowledge ID
func createKnowledgeBase(ip, port, apiKey, name string) (string, error) {
	// api endpoint
	endpoint := "api/v1/knowledge/create"

	payload := map[string]string{
		"name":        name,
		"description": "Customer Support Chat Logs",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	//req, err := http.NewRequest("POST", baseURL+"/knowledge/", bytes.NewBuffer(payloadBytes))
	url := fmt.Sprintf("http://%s:%s/%s", ip, port, endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to create knowledge base: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return "", fmt.Errorf("failed to decode error response: %w", err)
		}
		return "", fmt.Errorf("knowledge base creation failed with status %d: %s",
			resp.StatusCode, errResp.Detail)
	}

	var knowledgeResp KnowledgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&knowledgeResp); err != nil {
		return "", fmt.Errorf("*** failed to decode knowledge response: %w", err)
	}

	return knowledgeResp.ID, nil
}

// addFileToKnowledgeBase adds a file to a knowledge base
func addFileToKnowledgeBase(ip, port, apiKey, knowledgeID, fileID string) error {
	payload := map[string]string{"file_id": fileID}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := fmt.Sprintf("http://%s:%s", ip, port)
	req, err := http.NewRequest("POST",
		url+"/api/v1/knowledge/"+knowledgeID+"/file/add",
		bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to add file to knowledge base: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}
		return fmt.Errorf("adding file to knowledge base failed with status %d: %s",
			resp.StatusCode, errResp.Detail)
	}

	// if here, all is well
	return nil
}

func DoCreateKnowledgeBase(knowledgeBase string) (string, error) {
	var ip string
	var port string
	var apiKey string
	var err error
	var filePath string
	var errMsg string

	fmt.Println("Creating Knowledge base and seeding with sample file ... ")

	// get Ollama ip:port & token from env
	fmt.Print("  Getting environment variables\t\t\t: ")
	err = utils.GetOllamaEnvVars(&ip, &port, &apiKey)
	if err != nil {
		fmt.Println("ERROR")
		errMsg = fmt.Sprintf("Environemt error: %v", err)
		return "", errors.New(errMsg)
	}

	// get sample file
	err = utils.GetSampleKnowledgeFile(&filePath)
	if err != nil {
		fmt.Println("ERROR")
		errMsg = fmt.Sprintf("Sample Knowledge file error: %v", err)
		return "", errors.New(errMsg)
	}
	fmt.Println("OK")

	// Step 1: Upload the file
	fmt.Print("  Uploading sample file\t\t\t\t: ")
	fileID, err := uploadFile(ip, port, apiKey, filePath)
	if err != nil {
		fmt.Println("ERROR")
		errMsg = fmt.Sprintf("Error uploading file: %v", err)
		return "", errors.New(errMsg)
	}
	fmt.Println("OK")

	// Step 2: Create a knowledge base
	fmt.Print("  Creating Customer Support Knowledge Base\t: ")
	knowledgeID, err := createKnowledgeBase(ip, port, apiKey, knowledgeBase)
	if err != nil {
		fmt.Println("ERROR")
		errMsg = fmt.Sprintf("Error creating knowledge base: %v", err)
		return "", errors.New(errMsg)
	}
	fmt.Println("OK")

	// Step 3: Add the file to the knowledge base
	fmt.Print("  Adding uploaded file to new Knowledge Base\t: ")
	err = addFileToKnowledgeBase(ip, port, apiKey, knowledgeID, fileID)
	if err != nil {
		fmt.Println("ERROR")
		errMsg = fmt.Sprintf("Error adding file to knowledge base: %v\n", err)
		return "", errors.New(errMsg)
	}
	fmt.Println("OK")

	// if here, all is well
	return knowledgeID, nil
}
