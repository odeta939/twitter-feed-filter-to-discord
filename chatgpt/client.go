package chatgpt

import (
	"context"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)


func GetClient(config *Config) (*openai.Client) {
	client := openai.NewClient(config.AccessToken)
	return client;
}

	type Result struct {
		ShortSummary   string `json:"short_summary"`
		Sentiment string `json:"sentiment"`
	}


func EvaluateTweetSentiment(client *openai.Client, query string, ) (*Result, error) {

	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}
	
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Give a sentiment of this tweet with one word, as well as give a short summary of the tweet.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		
			ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "sentiment",
				Schema: schema,
				Strict: true,
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
	}
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
	}
	fmt.Println(result)
	return &result, nil

}