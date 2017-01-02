package main

import (
	"time"
	"strconv"
	"fmt"
)

type Config struct {
	Token            string
	Username         string
	LicenseFilename  string
	Branch           string
	CopyrightPattern string
	CommitMessage    string
	CurrentYear      string
	Repository       *RepositoryConfig
}

type RepositoryConfig struct {
	ListPerPage      int
	ListSortBy       string
	ListFilterByType string
}

type Labels struct {
	Start            string
	Token            string
	Username         string
	Required         string
	SelectAction     string
	LicenseFilename  string
	Branch           string
	CopyrightPattern string
	CommitMessage    string
	CurrentYear      string
	NextPageOrSkip   string
}

type Options struct {
	List     string
	Update   string
	Restart  string
	NextPage string
	Skip     string
}

var config = &Config{
	LicenseFilename: "LICENSE",
	Branch: "master",
	CopyrightPattern:  "Copyright \\(c\\) (\\d+)-(\\d+)",
	CommitMessage: "Update license.",
	CurrentYear: strconv.FormatInt(int64(time.Now().Year()), 10),
	Repository: &RepositoryConfig{
		ListPerPage: 30,
		ListSortBy: "name",
		ListFilterByType: "owner",
	},
}

var options = &Options{
	List: "l",
	Update: "u",
	Restart: "r",
	NextPage: "n",
	Skip: "s",
}

var labels = &Labels{
	Start: "\nQuit any time with Ctrl + C\n\n",
	Token: "\nGithub access token: ",
	Username: "Insert username: ",
	Required: "Required\n",
	SelectAction: fmt.Sprintf(
		"\nSelect your action: List repositories (%s = default) | Update license (%s) | Restart (%s): ",
		options.List,
		options.Update,
		options.Restart,
	),
	LicenseFilename: fmt.Sprintf("License filename (%s): ", config.LicenseFilename),
	Branch: fmt.Sprintf("Branch (%s): ", config.Branch),
	CopyrightPattern: fmt.Sprintf("Copyright pattern (%s): ", config.CopyrightPattern),
	CommitMessage: fmt.Sprintf("Commit message (%s): ", config.CommitMessage),
	CurrentYear: fmt.Sprintf("Current year (%s): ", config.CurrentYear),
	NextPageOrSkip: fmt.Sprintf("\nNext page (%s) | Skip (%s): ", options.NextPage, options.Skip),
}
