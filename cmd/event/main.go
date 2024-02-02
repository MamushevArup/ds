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
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	// Init bot instance
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err.Error())
	}

	b := bot.NewBot(session)
	err = b.StartBot()
	if err != nil {
		return
	}
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
