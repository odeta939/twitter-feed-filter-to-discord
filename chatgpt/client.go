package chatgpt

import (
	"context"
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
					Role:    openai.ChatMessageRoleSystem,
					Content: "Identify sentiment of the tweet with one of these words: (✅positive, ❌negative, 🔷neutral), as well as give a short summary of the tweet. For example: 'positive, this tweet is about a new product launch.'",
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
	return &result, nil

}