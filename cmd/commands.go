package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func StartApp() {
	commands := Commands{"Commands": "commands", "Insert": "insert", "Delete": "delete", "List": "list", "Exit": "exit"}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		command := scanner.Text()

		if command == commands["Exit"] {
			fmt.Println("Exiting the program...")
			break
		}
		if command == commands["Insert"] {
			fmt.Println("Insert the website that you want to track")
		}
		if command == commands["Delete"] {
			fmt.Println("Insert the website that you want to stop tracking")
		}
		if command == commands["List"] {
			fmt.Println("Listing all the tracked websites")
		}
		if command == commands["Commands"] {
			for _, value := range commands {
				fmt.Printf("%v\n", value)

			}
		}

	}
}
