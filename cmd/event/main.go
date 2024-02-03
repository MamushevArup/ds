package main

import (
	"fmt"
	"github.com/MamushevArup/discord-bot/cmd/event/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load environmental variables using 3 party
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	// Init bot instance
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err.Error())
	}

	// start new bot instance
	b := bot.NewBot(session)

	// implement message handling
	err = b.StartBot()
	if err != nil {
		return
	}
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
