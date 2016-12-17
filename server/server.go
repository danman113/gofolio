package server

import (
	"fmt"
	. "html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	. "github.com/danman113/gofolio/getgit"
)

// Todo:
// 	- Add logging module

var (
	templates *Template
	repos     *[]Repo
)

type RepoList struct {
	Repos []Repo
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		templates.Lookup("index.html").Execute(res, RepoList{*repos})
	} else {
		res.Header().Set("content-type", "text/html")
		fmt.Fprintf(res, "<h1>404</h1>")
	}
}

func view(res http.ResponseWriter, req *http.Request) {
	main := templates.Lookup("main.html")
	main.ExecuteTemplate(res, "main", struct{ Content string }{"Hello"})
}

// Sets up templates and starts server
func Run() {

	go func() {
		errors := setupTemplates("server/templates")
		if errors != nil {
			fmt.Println(errors)
		}
	}()

	fmt.Println("Starting server...")
	http.HandleFunc("/", index)
	http.HandleFunc("/view", view)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/static/"))))
	http.ListenAndServe(":8080", nil)
}

// Sets global repo list
func SetRepos(r *[]Repo) {
	repos = r
}

// Reads a directory and parses all of the templates
func setupTemplates(folder string) error {

	contents, err := ioutil.ReadDir(folder)

	if err != nil {
		return err
	}

	var files []string

	for _, file := range contents {
		full_name := file.Name()
		files = append(files, filepath.Join(folder, full_name))
	}

	var temperr error

	templates, temperr = ParseFiles(files...)

	if temperr != nil {
		return temperr
	}

	return nil
}
