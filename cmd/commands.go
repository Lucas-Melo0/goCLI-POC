package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func StartApp() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		command := scanner.Text()

		if command == "exit" {
			fmt.Println("Exiting the program...")
			break
		}

		if len(command) == 0 {
			fmt.Println("oiii")
		} else {
			fmt.Printf("You entered: %s\n", command)
		}
	}
}
