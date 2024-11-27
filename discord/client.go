package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetClient(config *Config) (*discordgo.Session, error) {
	if(config.Token == ""){
		return nil, fmt.Errorf("no token found")
	}

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord session: %w", err)
	}

	// TODO check if error returns the line of code where it failed
	err = discord.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}
	defer discord.Close()

	return discord , nil;
}

func SendMessage( channelID string, message string, discord *discordgo.Session) (error) {
	sentMessage, err := discord.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Println("Error sending message")
		return fmt.Errorf("failed to send message: %w", err)
	}
	fmt.Println("Message sent: ", sentMessage.Content)
	return nil
}