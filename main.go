package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Channel for the loading spinner
	done := make(chan bool)

	// Channel for the API response
	respChan := make(chan *genai.GenerateContentResponse, 1)
	errChan := make(chan error, 1)

	loadingSpinner := func(done chan bool) {
		spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0

		for {
			select {
			case <-done:
				fmt.Printf("\r%s\r", "                    ")
				return
			default:
				fmt.Printf("\r%s Thinking...", spinner[i])
				time.Sleep(100 * time.Millisecond)
				i = (i + 1) % len(spinner)
			}
		}
	}

	generateJoke := func() {
		model := client.GenerativeModel("gemini-2.0-flash")
		resp, err := model.GenerateContent(ctx, genai.Text("Tell a joke about AI"))
		if err != nil {
			errChan <- err
			return
		}
		respChan <- resp
	}

	// Start goroutines
	go loadingSpinner(done)
	go generateJoke()

	select {
	case resp := <-respChan:
		done <- true
		for _, candidate := range resp.Candidates {
			for _, part := range candidate.Content.Parts {
				fmt.Println(part)
			}
		}
	case err := <-errChan:
		done <- true
		log.Fatal(err)
	}
}
