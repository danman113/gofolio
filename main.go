package main

import (
	"fmt"
	"github.com/danman113/gofolio/getgit"
	"github.com/danman113/gofolio/server"
)

func main() {
	user := "danman113"
	repos, err := getgit.GetRepos(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Repos of %s:\n", user)
		for _, repo := range repos {
			fmt.Println(repo)
		}
		server.Run()
	}

}
