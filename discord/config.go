package discord

import (
	"fmt"
	"os"
)

type Config struct {
	Token     string
	ChannelID string
}

func LoadDiscordConfig() (*Config , error){
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil , fmt.Errorf("no token found")
	}
	channelID := os.Getenv("DISCORD_CHANNEL_ID")
	return &Config{Token: token, ChannelID: channelID} , nil
}