package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var token = "YOUR BOT TOKEN"
var botChID = "YOUR CHANNEL ID"

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
	go ParseRSS()
	time.Sleep(time.Second * 1)
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
