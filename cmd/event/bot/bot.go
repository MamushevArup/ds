package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	somethingWrong = "Something went wrong try again later"
)

type Bot struct {
	session *discordgo.Session
}

func NewBot(session *discordgo.Session) *Bot {
	return &Bot{
		session: session,
	}
}

// message responsible only for send message from server
type message struct {
	Message string `json:"message"`
}

// StartBot main entry point for `frontend of the bot`
func (b *Bot) StartBot() error {
	b.session.AddHandler(b.startBotMessage)
	err := b.session.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return err
	}
	return nil
}

// startBotMessage encapsulate the logic for command handling
func (b *Bot) startBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check if the message starts with the command prefix
	if strings.HasPrefix(m.Content, "!") {
		command := removeSlash(m.Content)
		split := strings.Split(command, " ")
		// Check the command and respond accordingly
		switch split[0] {
		case "hello":
			err := handleHello(s, m, split[0])
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		case "game":
			err := handleGame(s, m, command, split[0])
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		case "guess":
			err := guess(command, split[0], s, m)
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		case "help":
			err := handleHelp(split[0], s, m)
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		case "poll":
			err := handlePoll(command, split[0], s, m)
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		case "vote":
			// expected behavior here ex -> !vote -q Who is me -o Ron
			err := handleVote(command, split[0], s, m)
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		}
	}
}

// this function remove the firs character which means the command ex !, /
func removeSlash(command string) string {
	// ex !game -> game
	// !game 10 40 -> game 10 40
	return command[1:]
}

// utility function which is repeatable almost for all commands
func marshalling(response *http.Response, s *discordgo.Session, m *discordgo.MessageCreate) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Parse the JSON response
	var msg message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	// Send the message from the JSON response back to the Discord channel
	_, err = s.ChannelMessageSend(m.ChannelID, msg.Message)
	if err != nil {
		log.Printf("message is not sent due to %v", err)
		return err
	}
	return nil
}
