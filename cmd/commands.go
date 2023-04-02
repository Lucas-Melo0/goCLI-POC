package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"webcli/utils"
)

var TrackedWebsites = make([]string, 0)

func StartApp() {
	commands := Commands{"Commands": "commands", "Insert": "insert", "Delete": "delete", "List": "list", "Exit": "exit"}

	// Slice to store the tracked websites

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
			TrackedWebsites = append(TrackedWebsites, website)
			fmt.Printf("Website %s has been added to tracking\n", website)
			go checkWebsite(website)
		}
		if command == commands["Delete"] {
			if len(words) < 2 {
				fmt.Println("Please provide a website to delete")
				continue
			}
			websiteFound := false
			for i, v := range TrackedWebsites {
				if v == words[1] {
					TrackedWebsites = utils.Delete(TrackedWebsites, i)
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
			for i, website := range TrackedWebsites {
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

func checkWebsite(site string) {
	for range time.Tick(10 * time.Second) {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Printf("Error checking %s: %s\n", site, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Website %s is up and running!\n", site)
		} else {
			fmt.Printf("Website %s is down. Status code: %d\n", site, resp.StatusCode)
		}
	}
}
