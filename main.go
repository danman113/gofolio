package main

import (
	"flag"
	"fmt"

	"github.com/danman113/gofolio/getgit"
	"github.com/danman113/gofolio/server"
)

var (
	username string
	password string
)

func main() {
	flag.StringVar(&username, "username", "danman113", "Github Username")
	flag.StringVar(&password, "password", "", "Github Password")
	flag.Parse()

	repos, err := getgit.GetRepos(username, password)
	if err != nil {
		fmt.Println(err)
	} else {
		server.SetRepos(&repos)
		server.Run()
	}
}
