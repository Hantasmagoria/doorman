package main

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	// "net/http"
	// "os"
	// "os/signal"
	// "strings"
	// "syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	//create the discordgo session
	Session, _ = discordgo.New("Bot " + Token)
)

func init() {
	Session.Token = os.Getenv("BOT_TOKEN")
	// #COMMENTED CODE
	// if Session.Token == "" {
	// 	//handle empty discord token here
	// }
}

func main() {

	if Session.Token == "" {
		log.Println("You must provide a Discord Authentication Token.")
		return
	}

	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n ", err)
		os.Exit(1)
	}
	log.Printf(`Now running. Press CTRL-C to exit.`)
	//conenct to Discord
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	//exit
	Session.Close()
}
