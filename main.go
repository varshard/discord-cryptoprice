package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!price"

var token, secret string

func init() {
	token = os.Getenv("APP_TOKEN")
	secret = os.Getenv("SECRET")
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Discord connection error:", err)
		return
	}
	fmt.Println("Discord connected")

	dg.AddHandler(messageCreated)

	err = dg.Open()
	if err != nil {
		fmt.Println("Discord connection error:", err)
		return
	}

	// Waiting for termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Bye bye")
	dg.Close()
}

func messageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	inMsg := strings.ToLower(m.Content)
	if strings.HasPrefix(inMsg, prefix) {
		pair := strings.Trim(inMsg[len(prefix):], " ")

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %f", strings.ToUpper(pair), GetLastestPrice(pair)))
	}
}
