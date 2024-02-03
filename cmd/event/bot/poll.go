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
	Question string         `json:"question"`
	Options  map[int]string `json:"options"`
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
		return err
	}

	url := fmt.Sprintf("%s%s/%s", os.Getenv("SERVER"), split[0], m.Author.ID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if resp.StatusCode == http.StatusCreated {
		err = marshalling(resp, s, m)
		if err != nil {
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
	questionEndIndex := findOptionIndex(questionParts)
	if questionEndIndex == -1 {
		return nil, fmt.Errorf("Invalid command format. Missing options after -q.")
	}

	// Extract the question and options
	question := strings.Join(questionParts[:questionEndIndex], " ")
	options := make(map[int]string)

	for i := questionEndIndex; i < len(questionParts)-1; i += 2 {
		// Extract option number and corresponding option text
		optionText := questionParts[i+1]

		options[(i-questionEndIndex)/2] = optionText
	}
	return &poll{
		Question: question,
		Options:  options,
	}, nil
}
