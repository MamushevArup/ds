package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
)

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
	return marshalling(response, s, m)
}
