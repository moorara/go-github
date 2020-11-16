[![Go Doc][godoc-image]][godoc-url]
[![Build Status][workflow-image]][workflow-url]
[![Go Report Card][goreport-image]][goreport-url]
[![Test Coverage][coverage-image]][coverage-url]
[![Maintainability][maintainability-image]][maintainability-url]

# go-github

A simple REST client for [GitHub API v3](https://docs.github.com/rest).

## Quick Start

```go
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
```


[godoc-url]: https://pkg.go.dev/github.com/moorara/go-github
[godoc-image]: https://godoc.org/github.com/moorara/go-github?status.svg
[workflow-url]: https://github.com/moorara/go-github/actions
[workflow-image]: https://github.com/moorara/go-github/workflows/Main/badge.svg
[goreport-url]: https://goreportcard.com/report/github.com/moorara/go-github
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/go-github
[coverage-url]: https://codeclimate.com/github/moorara/go-github/test_coverage
[coverage-image]: https://api.codeclimate.com/v1/badges/3da9d932c98a5d9ce65e/test_coverage
[maintainability-url]: https://codeclimate.com/github/moorara/go-github/maintainability
[maintainability-image]: https://api.codeclimate.com/v1/badges/3da9d932c98a5d9ce65e/maintainability
