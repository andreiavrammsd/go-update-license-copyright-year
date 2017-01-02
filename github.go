package main

import (
	"github.com/google/go-github/github"
)

func NewGithubClient(token string) *github.Client {
	http := getHttpClient(token)
	return github.NewClient(http)
}
