package main

import (
	"context"
	"fmt"

	"github.com/moorara/go-github"
)

func main() {
	c := github.NewClient("")

	commits, resp, err := c.Repo("octocat", "Hello-World").Commits(context.Background(), 50, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pages: %+v\n", resp.Pages)
	fmt.Printf("Rate: %+v\n\n", resp.Rate)
	for _, c := range commits {
		fmt.Printf("%s\n", c.SHA)
	}
}
