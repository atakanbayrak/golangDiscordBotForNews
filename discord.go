package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token = "MTA1MjMxOTE5OTM3ODI3NjM5Mg.GSwY4b.Q6KrTZQBtnL4VJIAzBKd264dnTvURQ3Ewqz-cw"
var botChID = "1052320073391558696"

var Dg *discordgo.Session

func initSession() {
	//Bot token kullanarak yeni bir discord oturumu oluşturur
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("DC Oturumu Acilamadi. Error", err)
		l.Fatalf("DC Oturumu Acilamadi. Error", err)
	}

	Dg = dg
}

func ConnectToDC() {
	initSession()

	//Register messageCrate function
	Dg.AddHandler(messageCreate)

	Dg.Identify.Intents = discordgo.IntentGuildMessages

	// wait until CTRL+C
	fmt.Println("Bot calisiyor. Cikmak icin CTRL+C basiniz")
	l.Println("[INFO] RSSParserBot calisiyor.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Dg.Close()
	l.Println("[INFO] Bot durduruldu.")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!rssbot" {
		s.ChannelMessageSend(botChID, "Burdayim!")
		l.Printf("[INFO] %s Kullanicisi beni cagirdi", m.Author)

	}
}
