package main

import (
	"fmt"
)

func ListAction(config *Config) {
	fmt.Println(config)
}

func UpdateAction(config *Config) {
	config.Token = readInput(labels.Token, config.Token, true)
	config.LicenseFilename = readInput(labels.LicenseFilename, config.LicenseFilename, false)
	config.Branch = readInput(labels.Branch, config.Branch, false)
	config.CopyrightPattern = readInput(labels.CopyrightPattern, config.CopyrightPattern, false)
	config.CommitMessage = readInput(labels.CommitMessage, config.CommitMessage, false)
	config.CurrentYear = readInput(labels.CurrentYear, config.CurrentYear, false)
	fmt.Println(config)
}
