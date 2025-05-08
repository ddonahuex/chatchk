package open_webui

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"nethopper.io/utils"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type File struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Files    []File    `json:"files,omitempty"`
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"` // Added to capture error details
}

// Endpoint: POST /api/chat/completions
//
//	OpenAI API compatible chat comdpletion endpoint.
func DoChatCompletion(model string, prompt string, knowledgeBase string) error {
	// api endpoint
	endpoint := "api/chat/completions"

	fmt.Println("knowledgeBaseID: " + knowledgeBase)

	// locals
	var ip string
	var port string
	var apiKey string
	var err error
	var errMsg string

	// get Ollama ip:port & token from env
	err = utils.GetOllamaEnvVars(&ip, &port, &apiKey)
	if err != nil {
		return err
	}

	// URL for the Ollama API (updated to include /ollama prefix)
	url := fmt.Sprintf("http://%s:%s/%s", ip, port, endpoint)

	// Create the request payload

	requestBody := ChatRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Files: []File{
			{
				Type: "collection",
				ID:   knowledgeBase,
			},
		},
	}

	// JSON payload
	// Marshal the payload to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		errMsg = fmt.Sprintf("Error marshaling JSON: %v", err)
		return errors.New(errMsg)
	}

	// Log the exact JSON payload for debugging
	fmt.Printf("Request Payload: %s\n", string(jsonData))

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		errMsg = fmt.Sprintf("Error creating request: %v", err)
		return errors.New(errMsg)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("\nQuerying Open WebUI:\n")
	fmt.Printf("  Endpoint: %s\n", url)
	fmt.Printf("  Model: %s\n", model)
	fmt.Printf("  Prompt: '%s'\n\n", prompt)

	//fmt.Printf("http message:\n%s\n\n", bytes.NewBuffer(jsonData))
	fmt.Printf("Response: \n")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg = fmt.Sprintf("Error making request: %v", err)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg = fmt.Sprintf("Error reading response: %v", err)
		return errors.New(errMsg)
	}

	// Log the full response body for debugging
	fmt.Printf("Full Response Body: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		errMsg = fmt.Sprintf("Error: received status code %d\nResponse: %s", resp.StatusCode, string(body))
		return errors.New(errMsg)
	}

	var chatResp ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		errMsg = fmt.Sprintf("Error unmarshaling response: %v", err)
		return errors.New(errMsg)
	}

	// Check for API-level errors
	if chatResp.Error != nil {
		return fmt.Errorf("API error: %s", chatResp.Error.Message)
	}
	// Print the model's response
	if len(chatResp.Choices) > 0 {
		fmt.Println("Model response:", chatResp.Choices[0].Message.Content)
	} else {
		fmt.Println("No response choices found")
	}

	// if here, all is well
	return nil
}
