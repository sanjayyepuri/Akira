package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sanjayyepuri/Akira/router"
	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token string
	Debug bool
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.BoolVar(&Debug, "debug", false, "Debug level")
	flag.Parse()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go -t [BOT_TOKEN]")
		return
	}

	if Debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Token is " + Token)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Error("error creating Discord session,", err)
		return
	}

	commandRouter := router.NewRouter().WithPrefix("~")

	commandRouter.RegisterCommand("pog", handlePog)
	commandRouter.RegisterCommand("ping", handlePing)
	commandRouter.RegisterCommand("pong", handlePong)

	// Register the messageCreate func as a callback for MessageCreate events.
	//dg.AddHandler(messageCreate)
	dg.AddHandler(commandRouter.Handler)

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

func handlePog(request *discordgo.MessageCreate, response *discordgo.Session) {
	response.ChannelMessageSend(request.ChannelID, "POGGGGG!!!!")
}

func handlePing(request *discordgo.MessageCreate, response *discordgo.Session) {
	response.ChannelMessageSend(request.ChannelID, "Pong!")
}

func handlePong(request *discordgo.MessageCreate, response *discordgo.Session) {
	response.ChannelMessageSend(request.ChannelID, "Ping!")
}
