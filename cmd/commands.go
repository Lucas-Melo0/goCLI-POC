package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"webcli/utils"
)

func StartApp() {
	commands := Commands{"Commands": "commands", "Insert": "insert", "Delete": "delete", "List": "list", "Exit": "exit"}

	// Slice to store the tracked websites
	trackedWebsites := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		words := strings.Fields(input)

		if len(words) == 0 {
			continue
		}

		command := words[0]

		if command == commands["Exit"] {
			fmt.Println("Exiting the program...")
			break
		}
		if command == commands["Insert"] {
			if len(words) < 2 {
				fmt.Println("Please provide a website to track")
				continue
			}
			website := words[1]
			trackedWebsites = append(trackedWebsites, website)
			fmt.Printf("Website %s has been added to tracking\n", website)
		}
		if command == commands["Delete"] {
			if len(words) < 2 {
				fmt.Println("Please provide a website to delete")
				continue
			}
			websiteFound := false
			for i, v := range trackedWebsites {
				if v == words[1] {
					trackedWebsites = utils.Delete(trackedWebsites, i)
					websiteFound = true
					break
				}
			}
			if !websiteFound {
				fmt.Println("We could not find this website in the list")
			} else {
				fmt.Println("Website deleted")
			}
		}
		if command == commands["List"] {
			fmt.Println("Listing all the tracked websites")
			for i, website := range trackedWebsites {
				fmt.Printf("%d: %s\n", i+1, website)
			}
		}
		if command == commands["Commands"] {
			for _, value := range commands {
				fmt.Printf("%v\n", value)
			}
		}
	}
}
