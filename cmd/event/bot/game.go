package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func handleGame(s *discordgo.Session, m *discordgo.MessageCreate, command, split string) error {
	low, high, err := parseGameCommand(command)
	if err != nil {
		fmt.Println("Error parsing game command:", err)
		return err
	}

	// Create a request body with the boundaries
	payload := map[string]interface{}{
		"lower":   low,
		"upper":   high,
		"user_id": m.Author.ID, // Replace with the actual Discord user ID
	}
	// Convert the map to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON payload:", err)
		return err
	}

	// Make an HTTP POST request to the "/game" endpoint
	url := fmt.Sprintf("%s%s", os.Getenv("SERVER"), split) // Replace with the actual endpoint
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return err
	}
	defer resp.Body.Close()
	// Check the response
	return marshalling(resp, s, m)
}

func guess(command, split string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	gs := strings.Fields(command)
	number, err := parseInt(gs[1])
	if err != nil {
		return err
	}
	// http://host/guess/:id/:number
	url := fmt.Sprintf("%s%s/%s/%d", os.Getenv("SERVER"), split, m.Author.ID, number)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error with sending request ", err)
		return err
	}
	defer resp.Body.Close()
	err = marshalling(resp, s, m)
	if err != nil {
		return err
	}
	return nil
}

func parseGameCommand(command string) (int, int, error) {
	// Split the command by spaces
	parts := strings.Fields(command)

	// Check if the command has at least three parts (including the "!game" command)
	if len(parts) < 3 {
		return 0, 0, fmt.Errorf("invalid !game command format")
	}

	// Parse the low and high boundaries
	low, err := parseInt(parts[1])
	if err != nil {
		return 0, 0, err
	}

	high, err := parseInt(parts[2])
	if err != nil {
		return 0, 0, err
	}

	return low, high, nil
}

func parseInt(s string) (int, error) {
	// Parse the string to an integer
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %s", s)
	}
	return val, nil
}
