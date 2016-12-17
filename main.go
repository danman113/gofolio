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
		server.SetRepos(&repos)
		server.Run()
	}
}
