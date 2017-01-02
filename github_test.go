package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/google/go-github/github"
)

func TestNewGithubClientWithToken(t *testing.T) {
	client := NewGithubClient("token")
	assert.IsType(t, client, &github.Client{})
}

func TestNewGithubClientWithOutToken(t *testing.T) {
	client := NewGithubClient("")
	assert.IsType(t, client, &github.Client{})
}
