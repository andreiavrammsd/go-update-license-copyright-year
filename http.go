package main

import (
	"golang.org/x/oauth2"
	"net/http"
)

func getHttpClient(token string) *http.Client {
	var ts oauth2.TokenSource
	if len(token) > 0 {
		ts = oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: token,
			},
		)
	} else {
		ts = nil
	}

	return oauth2.NewClient(oauth2.NoContext, ts)
}
