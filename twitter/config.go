package twitter

import (
	"fmt"
	"os"
)

type Config struct {
	BearerToken string
}

func LoadTwitterConfig() (*Config, error) {
	bearerToken := os.Getenv("GOTWI_BEARER_TOKEN")
	if bearerToken == "" {
		return nil, fmt.Errorf("no token found")
	}
	return &Config{BearerToken: bearerToken}, nil
}