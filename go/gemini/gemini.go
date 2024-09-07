// Package gemini provides functionality for filtering social media posts.
package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

/*
Filters social media posts that are not relevant to a job search and title.
*/
func FilterSocialMediaPosts(inputPosts []string, title string) ([]string, error) {
	placeHolder := `
		You are a social media post filter. I am a %s who is looking for a job. 
		I want to remove social media posts that are not relevant to my job search and title.
		When providing a response ONLY return a JSON list of the relevant social media posts.
		DO NOT return code blocks or any other information, ONLY the JSON.
		Below are the social media posts that I would like you to filter:

	`
	// Generate the prompt
	prompt := fmt.Sprintf(placeHolder, title)

	// Append the posts to the prompt
	for _, post := range inputPosts {
		prompt += "> " + post + "\n"
	}

	// Generate the content
	res, err := generate(prompt)
	if err != nil {
		return []string{}, err
	}

	// convert from bytes to string array
	var outputPosts []string
	err = json.Unmarshal(res, &outputPosts)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return []string{}, err
	}

	return outputPosts, nil
}

/*
Generate generates content using the provided input string
*/
func generate(input string) ([]byte, error) {
	ctx := context.Background()

	client := getClient(ctx)
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return []byte{}, err
	}

	return getResponseBytes(resp), nil
}

/*
Get a new instance of the genai.Client with the provided API key.
It takes a context.Context as a parameter and returns a pointer to the genai.Client.
If an error occurs while creating the client, it logs the error and exits the program.
*/
func getClient(ctx context.Context) *genai.Client {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

/*
Get response bytes from the genai.GenerateContentResponse
*/
func getResponseBytes(resp *genai.GenerateContentResponse) []byte {
	var buffer bytes.Buffer
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Fprintf(&buffer, "%s", part)
			}
		}
	}
	return buffer.Bytes()
}
