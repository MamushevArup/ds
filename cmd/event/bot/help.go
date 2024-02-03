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

func handleHelp(split string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	url := fmt.Sprintf("%s%s", os.Getenv("SERVER"), split)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	var hm map[string]string
	err = json.Unmarshal(byteValue, &hm)
	if err != nil {
		log.Println(err)
		return err
	}
	res := formatJSON(hm)
	_, err = s.ChannelMessageSend(m.ChannelID, res)
	if err != nil {
		log.Printf("message is not sent due to %v", err)
		return err
	}
	return nil
}

func formatJSON(data map[string]string) string {
	var formattedStrings []string

	for key, value := range data {
		formattedStrings = append(formattedStrings, fmt.Sprintf("%s : %v", key, value))
	}

	return strings.Join(formattedStrings, "\n")
}
