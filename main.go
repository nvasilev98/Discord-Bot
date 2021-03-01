package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

const token string = "ODE1NTA1OTA0MTcwNDM0NTYw.YDtZDg.HTp9ykyMKhCRmA8x1FfwNUDFFkA"

var t0 time.Time
var BotID string

func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	id, err := dg.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	BotID = id.ID
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if s == nil || m == nil {
		return
	}

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "" {
		return
	}
	if m.Content == "alkoholik li sum" {
		s.ChannelMessageSend(m.ChannelID, "DA")
	}

	if m.Content[0] == '!' && strings.Count(m.Content, "!") < 2 {

		ExecuteCommand(s, m.Message, t0)
		return
	}

}

func HandlePlayCommand(s *discordgo.Session, game string) {
	err := s.UpdateGameStatus(0, game)
	if err != nil {
		println("[Error] Issue while updating bot status: ", err)
		return
	}
}
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {

	msg := strings.Split(strings.TrimSpace(m.Content), "!")[1]

	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], "!")[1]
	}

	switch msg {
	case "play":
		game := strings.Join(strings.Split(m.Content, " ")[1:], " ")
		HandlePlayCommand(s, game)
	}
}
