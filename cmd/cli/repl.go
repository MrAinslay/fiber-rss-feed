package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MrAinslay/fiber-rss-feed/internal/api"
)

type ApiConfig struct {
	ApiClient api.Client
	ApiKey    string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *ApiConfig, s string) error
}

func cleanInput(s string) string {
	output := strings.TrimSpace(s)
	output = strings.ToLower(output)
	return output
}

func printPrompt() {
	log.Print("Rss-Feed > ")
}

func StartRepl(cfg *ApiConfig) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		splitText := strings.Split(text, " ")
		if command, exists := commands[splitText[0]]; exists {
			if len(splitText) > 1 {
				err := command.callback(cfg, fmt.Sprint(splitText[1:]))
				if err != nil {
					log.Println(err)
				}
			} else {
				err := command.callback(cfg, splitText[0])
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			log.Println("Unknown Command")
		}
		printPrompt()
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"create": {
			name:        "create",
			description: "Creates either a user or a feed depending on the selected option",
			callback:    commandCreate,
		},
	}
}
