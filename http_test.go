package main

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetHttpClientWithToken(t *testing.T) {
	client := getHttpClient("token")
	assert.IsType(t, client, &http.Client{})
}

func TestGetHttpClientWithoutToken(t *testing.T) {
	client := getHttpClient("")
	assert.IsType(t, client, &http.Client{})
}
