package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
	"strings"
)

// !poll <question> <option>
type poll struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

func handlePoll(command, split string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	pollData, err := parsePollCommand(command)
	if err != nil {
		fmt.Println("Error parsing poll command:", err)
		_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}
	// Marshal the data into JSON
	jsonData, err := json.Marshal(pollData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}

	url := fmt.Sprintf("%s%s/%s", os.Getenv("SERVER"), split, m.Author.ID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}
	if resp.StatusCode == http.StatusCreated {
		err = marshalling(resp, s, m)
		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
			return err
		}
	}
	return nil
}

func parsePollCommand(command string) (*poll, error) {

	parts := strings.Fields(command)
	if len(parts) < 4 || parts[1] != "-q" {
		return nil, fmt.Errorf("Invalid command format. Usage: !poll -q <question> -1 <option1> -2 <option2> ...")
	}

	questionParts := parts[2:]
	questionEndIndex := findPollIndex(questionParts)
	if questionEndIndex == -1 {
		return nil, fmt.Errorf("Invalid command format. Missing options after -q.")
	}

	// Extract the question and options
	question := strings.Join(questionParts[:questionEndIndex], " ")
	options := make([]string, 0)

	for i := questionEndIndex; i < len(questionParts)-1; i += 2 {
		// Extract option number and corresponding option text
		optionText := questionParts[i+1]

		options = append(options, optionText)
	}
	return &poll{
		Question: question,
		Options:  options,
	}, nil
}

func findPollIndex(parts []string) int {
	for i, part := range parts {
		if strings.HasPrefix(part, "-") && i+1 < len(parts) {
			return i
		}
	}
	return -1
}

func handleVote(command, split string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	v, err := parseVote(command)
	if err != nil {
		_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}
	url := fmt.Sprintf("%s%s/%s", os.Getenv("SERVER"), split, m.Author.ID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}
	if resp.StatusCode == http.StatusCreated {
		err = marshalling(resp, s, m)
		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
			return err
		}
	}
	return nil
}

type vote struct {
	Question string `json:"question"`
	Option   string `json:"option"`
}

func parseVote(command string) (*vote, error) {
	parts := strings.Fields(command)

	questionParts := parts[2:]
	questionEndIndex := findQuestionIndex(questionParts)
	if questionEndIndex == -1 {
		return nil, fmt.Errorf("Invalid command format. Missing options after -q.")
	}

	question := strings.Join(questionParts[:questionEndIndex], " ")
	option := strings.Join(questionParts[questionEndIndex+1:], " ")
	// replace with actual user ID
	return &vote{
		Question: question,
		Option:   option,
	}, nil
}

func findQuestionIndex(parts []string) int {
	for i, part := range parts {
		if strings.HasPrefix(part, "-") && (part == "-q" || part == "-o") && i+1 < len(parts) {
			return i
		}
	}
	return -1
}
