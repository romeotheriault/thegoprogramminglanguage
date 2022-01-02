// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// Sort the issues by CreatedAt time.
	sort.Slice(result.Items, func(i, j int) bool { return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt) })

	now := time.Now()
	month := now.Add(-24 * 30 * time.Hour)
	year := now.Add(-24 * 365 * time.Hour)

	ms, ys, ya := 0, 0, 0

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		switch {
		case month.Before(item.CreatedAt):
			if ms == 0 {
				ms = 1
				fmt.Printf("-------------- In the last month --------------\n")
			}
			fmt.Printf("#%-5d %9.9s %.75s\n", item.Number, item.User.Login, item.Title)
		case year.Before(item.CreatedAt):
			if ys == 0 {
				ys = 1
				fmt.Printf("-------------- In the last year --------------\n")
			}
			fmt.Printf("#%-5d %9.9s %.75s\n", item.Number, item.User.Login, item.Title)
		case year.After(item.CreatedAt):
			if ya == 0 {
				ya = 1
				fmt.Printf("-------------- Over a year ago --------------\n")
			}
			fmt.Printf("#%-5d %9.9s %.75s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

//$ go build gopl.io/ch4/issues
//$ ./issues repo:golang/go is:open json decoder
