package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"nethopper.io/admin"
	"nethopper.io/ingest"
	"nethopper.io/knowledge"
	"nethopper.io/ollama"
	"nethopper.io/open_webui"
	"nethopper.io/prompts"
)

func main() {
	// locals
	var err error
	var prompt string
	var knowledgeBase string
	var knowledgeID string
	var knowledgeCreated bool = false
	var choice int

	// defaults
	model := "gemma2:9b"

	// nothing burgers (for now)
	ingest.DoSomething()
	prompts.DoSomething()
	admin.DoSomething()

	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		choice, err = getUserChoice(reader)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			prompt = "Why is the sky blue?"
			err = open_webui.DoChatCompletion(model, prompt, "")
			if err != nil {
				fmt.Println("Open WebUI Error: ", err)
			}
		case 2:
			prompt = "Why is the sky blue?"
			err = ollama.DoGenerate(model, prompt)
			if err != nil {
				fmt.Println("Ollama Error: ", err)
			}
		case 3:
			// only create once
			if !knowledgeCreated {
				knowledgeBase = "Acme Customer Support Chats"
				knowledgeID, err = knowledge.DoCreateKnowledgeBase(knowledgeBase)
				if err != nil {
					fmt.Println("Knowledge Base Error: ", err)
					continue
				}
				fmt.Println("The knowledgeID is " + knowledgeID)
				knowledgeCreated = true
			} else {
				fmt.Println("Knowledge Base already successfully created")
			}
		case 4:
			prompt = "Summarize the customer problem and how it was resolved. Be verbose."
			err = open_webui.DoChatCompletion(model, prompt, knowledgeID)
			if err != nil {
				fmt.Println("Open WebUI Error: ", err)
			}
			return
		case 5:
			fmt.Println("\nExiting chatchk.")
			return
		default:
			fmt.Println("Invalid option. Please select a number between 1 and 5.")
		}
	}
}

func displayMenu() {
	fmt.Println("\n\n============ Chatchk Menu ============")
	fmt.Println("1. Inference example - Open WebUI")
	fmt.Println("2. Inference example - Ollama")
	fmt.Println("3. Create Knowledge Base - Open WebUI")
	fmt.Println("4. RAG Knowledge Base - Open WebUI")
	fmt.Println("5. Exit")

	fmt.Print("\nSelect an option (1-5): ")
}

func getUserChoice(reader *bufio.Reader) (int, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return choice, nil
}
