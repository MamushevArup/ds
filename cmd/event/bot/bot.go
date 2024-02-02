package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var somethingWrong = "Something went wrong try again later"

type Bot struct {
	session *discordgo.Session
	// this field handle last command typed
	last []string
}

func NewBot(session *discordgo.Session) *Bot {
	return &Bot{
		session: session,
		last:    make([]string, 0),
	}
}

type message struct {
	Message string `json:"message"`
}

func (b *Bot) StartBot() error {
	b.session.AddHandler(b.startBotMessage)
	err := b.session.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return err
	}
	return nil
}

func (b *Bot) startBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check if the message starts with the command prefix
	if strings.HasPrefix(m.Content, "!") {
		command := removeSlash(m.Content)
		// Check the command and respond accordingly
		switch command {
		case "hello":
			err := handleHello(s, m, command)
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, somethingWrong)
				return
			}
		case "game":

		}
	}
	b.last = append(b.last, m.Content)
	fmt.Println(b.last)
}

func handleHello(s *discordgo.Session, m *discordgo.MessageCreate, command string) error {
	// the request should look like this host/hello/id
	endpoint := fmt.Sprintf("%s%s/%s", os.Getenv("SERVER"), command, m.Author.ID)
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
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

// this function remove the firs character which means the command ex !, /
func removeSlash(command string) string {
	return command[1:]
}
