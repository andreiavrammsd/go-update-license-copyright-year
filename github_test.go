package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/google/go-github/github"
)

func TestNewGithubClientWithToken(t *testing.T) {
	config := &Config{
		Token: "token",
	}
	client := NewGithubClient(config)
	assert.IsType(t, client, &GithubClient{})
	assert.IsType(t, client.Client, &github.Client{})
	assert.Equal(t, config, client.Config)
}

func TestNewGithubClientWithOutToken(t *testing.T) {
	config = &Config{}
	client := NewGithubClient(config)
	assert.IsType(t, client, &GithubClient{})
	assert.IsType(t, client.Client, &github.Client{})
	assert.Equal(t, config, client.Config)
}
