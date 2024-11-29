package chatgpt

import (
	"fmt"
	"os"
)

type Config struct {
	AccessToken string
}

func LoadConfig() (*Config, error) { 
	accessToken := os.Getenv("OPENAI_API_KEY")
	if accessToken == "" {
		return nil, fmt.Errorf("no token found")
	}
	return &Config{AccessToken: accessToken}, nil
}