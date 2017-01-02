package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

var scan = func() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func readInput(text, defaultValue string, required bool) string {
	for {
		fmt.Print(text)
		input := strings.Trim(scan(), " ")

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
