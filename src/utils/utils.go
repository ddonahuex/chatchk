package utils

import (
	"errors"
	"os"
)

// GetOllamaEnvVars
//   - Sets input strings to Ollama env vars for IP address, Port, and api key
//   - Returns an error string (if applic.)
func GetOllamaEnvVars(ip *string, port *string, apiKey *string) error {
	var value string
	var exists bool

	value, exists = os.LookupEnv("OLLAMA_IP")
	if !exists {
		return errors.New("OLLAMA_IP environment variable not set")
	}
	*ip = value

	value, exists = os.LookupEnv("OLLAMA_PORT")
	if !exists {
		return errors.New("OLLAMA_PORT environment variable not set")
	}
	*port = value

	value, exists = os.LookupEnv("OLLAMA_API_KEY")
	if !exists {
		return errors.New("OLLAMA_API_KEY environment variable not set")
	}
	*apiKey = value

	// if here, all is well
	return nil
}

func GetSampleKnowledgeFile(filePath *string) error {
	var value string
	var exists bool

	value, exists = os.LookupEnv("OLLAMA_KNOWLEDGE_FILE")
	if !exists {
		return errors.New("OLLAMA_KNOWLEDGE_FILE environment variable not set")
	}
	*filePath = value

	// if here, all is well
	return nil
}
