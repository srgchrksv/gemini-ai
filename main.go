package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// readLine users input prompt
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {
	// init environment variable
	API_KEY := flag.String("api-key", os.Getenv("API_KEY"), "Gemini AI api key required")
	flag.Parse()
	if *API_KEY == "" {
		log.Fatal("api key required")
	}

	// init genai client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(*API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// init model
	model := client.GenerativeModel("gemini-1.0-pro")
	// start chat with session history
	cs := model.StartChat()

	// chat
	for {
		fmt.Println("User prompt:")
		input, err := readLine()
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cs.SendMessage(ctx, genai.Text(input))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("*****************\nGemini Response: %v \n*****************\n", resp.Candidates[0].Content.Parts[0])

		fmt.Println("Wish to continue?  y/n and press enter")
		input, err = readLine()
		if err != nil {
			log.Fatal(err)
		}
		if input == "n" {
			break
		}
		fmt.Println("==================")
	}

}
