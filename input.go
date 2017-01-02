package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func readInput(text, defaultValue string, required bool) string {
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print(text)
		scanner.Scan()
		input := strings.Trim(scanner.Text(), " ")

		if required == true {
			if len(input) == 0 {
				fmt.Print(labels.Required)
				continue
			}

			return input
		} else {
			if len(input) > 0 {
				return input
			}

			return defaultValue
		}
	}
}
