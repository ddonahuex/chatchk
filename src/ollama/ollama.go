package ollama

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"nethopper.io/utils"
)

// RunQuery: Use Ollama API to execute a given prompt using a given model
func DoGenerate(model string, prompt string) error {
	// api endpoint
	endpoint := "api/generate"

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
	url := fmt.Sprintf("http://%s:%s/ollama/%s", ip, port, endpoint)

	// JSON payload
	payload := []byte(fmt.Sprintf(`{"model": "%s", "prompt": "%s"}`, model, prompt))

	fmt.Printf("\nQuerying ollama:\n")
	fmt.Printf("  Endpoint: %s\n", url)
	fmt.Printf("  Model: %s\n", model)
	fmt.Printf("  Prompt: '%s'\n\n", prompt)
	fmt.Printf("Response: \n")

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		errMsg = fmt.Sprintf("Error creating request: %v\n", err)
		return errors.New(errMsg)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		errMsg = fmt.Sprintf("Error making request: %v\n", err)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg = fmt.Sprintf("Error reading response: %v\n", err)
		return errors.New(errMsg)
	}

	// Print the response
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", string(body))

	// if here, all is well
	return nil
}
