package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/michimani/gotwi"
	"github.com/odeta939/twitter-feed-filter-to-discord/chatgpt"
	"github.com/odeta939/twitter-feed-filter-to-discord/discord"
	"github.com/odeta939/twitter-feed-filter-to-discord/twitter"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

func main() {
	//** ChatGPT **//
	chatGPTConfig, err := chatgpt.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load ChatGPT configuration: %v", err)
	}

	openaiClient := chatgpt.GetClient(chatGPTConfig)

	

	//** Discord **//
	discoedConfig, err := discord.LoadDiscordConfig()
	if err != nil {
		log.Fatalf("Failed to load Discord configuration: %v", err)
	}

	discordClient, err := discord.GetClient(discoedConfig)
	if err != nil {
		log.Fatalf("Failed to create Discord client: %v", err)
	}
	
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

	tweets, err := twitter.GetRecentTweets(twitterClient, "#authjs OR #nextauth")
	if err != nil {
		log.Fatalf("Failed to get recent tweets: %v\n", err)
	}

	for _, tweet := range tweets {
		
		evaluation, err := chatgpt.EvaluateTweetSentiment(openaiClient, gotwi.StringValue(tweet.Text))
		if err != nil {
			log.Fatalf("Failed to evaluate tweet sentiment: %v\n", err)
		}

		discordMessage := GenerateDiscordMessage(*evaluation, twitter.GetTweetUrl(tweet))

		discord.SendMessage(discoedConfig.ChannelID, discordMessage, discordClient)
	}

	defer discordClient.Close()

}

func GenerateDiscordMessage(evaluation chatgpt.Result, tweetUrl string ) string {
	return fmt.Sprintf("**%s**\n%s\n %s", evaluation.Sentiment, evaluation.ShortSummary, tweetUrl)
}