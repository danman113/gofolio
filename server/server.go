package server

import (
	"fmt"
	. "html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// Todo: 
// 	- Add logging module

var (
	templates map[string]*Template = make(map[string]*Template)
)

func setupTemplates(folder string) []error {
	var errors []error

	contents, err := ioutil.ReadDir(folder)

	if err != nil {
		return append(errors, err)
	}

	for _, file := range contents {
		full_name := file.Name()
		name := full_name[:len(full_name)-5]
		temp, temperr := ParseFiles(filepath.Join(folder, full_name))
		if temperr != nil {
			errors = append(errors, temperr)
			continue
		} else {
			fmt.Println("Added template " + name)
			templates[name] = temp
		}
	}

	return errors
}

func index(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(res, "<h1>Hello Vivian!</h1>")
	templates["index"].Execute(res, nil)
}

func Run() {

	go func () {
		errors := setupTemplates("server/templates")
		if len(errors) > 0 {
			for _, val := range errors {
				fmt.Println( val )
			}
		}
	}()
	
	fmt.Println("Starting server...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
