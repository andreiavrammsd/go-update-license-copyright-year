package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

func ListAction(config *Config) {
	client := NewGithubClient(config.Token)
	page := 1

	Loop:
	for ;; {
		o := &github.RepositoryListOptions{
			Sort: config.Repository.ListSortBy,
			Type: config.Repository.ListFilterByType,
			ListOptions: github.ListOptions{
				Page: page,
				PerPage: config.Repository.ListPerPage,
			},
		}
		repos, resp, err := client.Repositories.List(config.Username, o)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, repo := range repos {
			fmt.Printf("\n%s", *repo.Name)
		}

		if resp.NextPage == 0 {
			break Loop
		}

		option := readInput(labels.NextPageOrSkip, options.NextPage, false)
		switch option {
		case options.Skip:
			break Loop
		default:
			page++
		}
	}

	fmt.Println()
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
