package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/michimani/gotwi"
	"github.com/odeta939/twitter-feed-filter-to-discord/discord"
	"github.com/odeta939/twitter-feed-filter-to-discord/twitter"
)

func init() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

func main() {

	//** Discord **//
	discoedConfig, err := discord.LoadDiscordConfig()
	if err != nil {
		log.Fatalf("Failed to load Discord configuration: %v", err)
	}

	discordClient, err := discord.GetClient(discoedConfig)
	if err != nil {
		log.Fatalf("Failed to create Discord client: %v", err)
	}
	

	fmt.Println("Discord client successfully initialized!")


	//** Twitter **//
	twitterConfig, err := twitter.LoadTwitterConfig()
	if err != nil {
		log.Fatalf("Failed to load Twitter configuration: %v", err)
	}

	twitterClient, err := twitter.GetClient(twitterConfig)
	if err != nil {
		log.Fatalf("Failed to create Twitter client: %v", err)
	}

	fmt.Println("Twitter client successfully initialized!")

	res, err := twitter.GetRecentTweets(twitterClient, "#authjs")
	if err != nil {
		log.Fatalf("Failed to get recent tweets: %v", err)
	}

	for _, tweet := range res.Data {
		discord.SendMessage(discoedConfig.ChannelID,  gotwi.StringValue(tweet.Text), discordClient)
		fmt.Printf("Tweet: %s\n",  gotwi.StringValue(tweet.Text))
	}

	defer discordClient.Close()

}