package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"github-cli/github"
)

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func usage() {
	fmt.Printf("ghissues <create|read|update|close|list> -o owner -r repo [-t ticket#] [-T title]\n")
	os.Exit(1)
}

func main() {
	//var issue Issue

	var action string
	var owner string
	var repo string
	var ticket string
	var title string

	arglen := len(os.Args)
	if arglen < 6 {
		usage()
	}

	for i, v := range os.Args[1:] {
		if i == 0 {
			if v == "create" || v == "list" || v == "read" || v == "update" || v == "close" {
				action = v
			} else {
				usage()
			}
		}
		if v == "-o" {
			owner = os.Args[(i + 2)]
		}
		if v == "-r" {
			repo = os.Args[(i + 2)]
		}
		if v == "-t" {
			ticket = os.Args[(i + 2)]
		}
		if v == "-T" {
			title = os.Args[(i + 2)]
		}
	}

	// Make sure owner and repo are set.
	if owner == "" || repo == "" {
		usage()
	}

	repoUrl := "https://api.github.com/repos/" + owner + "/" + repo + "/issues"

	switch action {
	case "create":
		fmt.Println("create")
		if title == "" {
			usage()
		}
	case "list":
		resp, err := http.Get(repoUrl)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Printf("search query failed: %s", resp.Status)
		}
		var result []IssuesSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			fmt.Printf("json decoding failed: %s\n", err)
		}
		for _, i := range result {
			for _, item := range i.Items {
				fmt.Printf("#%-5d %9.9s %.75s\n", item.Number, item.User.Login, item.Title)
			}
		}
		resp.Body.Close()
	case "read":
		fmt.Println("read")
		if ticket == "" {
			usage()
		}
	case "update":
		fmt.Println("update")
		if ticket == "" {
			usage()
		}
	case "close":
		fmt.Println("close")
		if ticket == "" {
			usage()
		}
	}
}
