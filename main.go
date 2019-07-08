package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go -t [BOT_TOKEN]")
		return
	}

	log.Info("Token is " + Token)

	// Create a new Discord session using the provided bot token.

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Error("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.Bot {
		log.Infof("ignoring author %s since it is a bot", m.Author.ID)
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content != "" && m.Content[0] == '~' {
		command := m.Content[1:]

		log.Infof("[COMMAND]: %s", command)

		if command == "ping" {
			s.ChannelMessageSend(m.ChannelID, "Pog!")
		}

		// If the message is "pong" reply with "Ping!"
		if command == "pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}
	}
}
