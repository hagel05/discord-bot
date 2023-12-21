package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Token      string
	Command    string
	ChannelID  string
	TextToSent string
)

func main() {
	flag.StringVar(&Token, "t", os.Getenv("DISCORD_BOT_TOKEN"), "Bot Token")
	flag.StringVar(&Command, "command", "sendtext", "Command to execute (e.g., sendtext)")
	flag.StringVar(&ChannelID, "channel", os.Getenv("DISCORD_CHANNEL_ID"), "Channel ID")
	flag.StringVar(&TextToSent, "text", "", "Text to send")

	flag.Parse()

	if TextToSent == "" {
		fmt.Println("Usage: go run main.go -t YOUR_BOT_TOKEN -command sendtext -channel CHANNEL_ID -text 'Text to send'")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		os.Exit(1)
	}

	//dg.AddMessageCreate(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session,", err)
		os.Exit(1)
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	switch strings.ToLower(Command) {
	case "sendtext":
		err := sendMessageToChannel(dg, ChannelID, TextToSent)
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		} else {
			fmt.Printf("Text sent to channel %s\n", ChannelID)
		}
	default:
		fmt.Println("Unknown command. Supported commands: sendtext")
	}

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Handle incoming messages if needed
}

func sendMessageToChannel(s *discordgo.Session, channelID, text string) error {
	_, err := s.ChannelMessageSend(channelID, text)
	return err
}
