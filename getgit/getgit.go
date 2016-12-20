package getgit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type JSONObject map[string]interface{}
type JSONArray []map[string]interface{}

var (
	// Auth info for github API
	username string
	password string

	// Global client
	client *http.Client = &http.Client{}
)

// Simplifies getting request
func getHTTP(url string) ([]byte, error) {

	req, rerr := http.NewRequest("GET", url, nil)
	if rerr != nil {
		return nil, rerr
	}

	if password != "" {
		req.SetBasicAuth(username, password)
	}

	res, geterr := client.Do(req)

	if geterr != nil {
		return nil, geterr
	}
	defer res.Body.Close()

	val, readerr := ioutil.ReadAll(res.Body)
	if readerr != nil {
		return nil, readerr
	}
	return val, nil
}

// Tests an interface for nil. If so, returns the null string
func NilString(n interface{}) string {

	if n == nil {
		return ""
	} else {
		return n.(string)
	}
}

// Parses proper github JSON to a Repo object
func ParseRepo(obj JSONObject) Repo {

	id := int(obj["id"].(float64))
	name := obj["name"].(string)
	full_name := NilString(obj["full_name"])
	url := NilString(obj["html_url"])
	description := NilString(obj["description"])
	created, _ := time.Parse(time.RFC3339, obj["created_at"].(string))
	updated, _ := time.Parse(time.RFC3339, obj["updated_at"].(string))
	homepage := NilString(obj["homepage"])
	stars := int(obj["stargazers_count"].(float64))
	watchers := int(obj["watchers_count"].(float64))
	language := NilString(obj["language"])
	git_url := NilString(obj["git_url"])
	contents_url := NilString(obj["contents_url"])
	contents_url = contents_url[:len(contents_url)-7]
	return Repo{
		Id:           id,
		Name:         name,
		Full_name:    full_name,
		URL:          url,
		Description:  description,
		Created:      created,
		Updated:      updated,
		Homepage:     homepage,
		Stars:        stars,
		Watchers:     watchers,
		Language:     language,
		Git_URL:      git_url,
		Contents_URL: contents_url,
	}
}

func GetContentData(repo *Repo) {
	cURL := repo.Contents_URL
	fmt.Println(cURL)

	val, err := getHTTP(cURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	var rawContents JSONArray
	jsonerr := json.Unmarshal(val, &rawContents)
	if jsonerr != nil {
		fmt.Println(jsonerr)
		return
	}

	for _, file := range rawContents {
		FileSearch(repo, file)
	}

}

func FileSearch(repo *Repo, file JSONObject) {
	filename := file["name"].(string)
	var imageName string
	if len(filename) > 4 && strings.Contains(filename, ".") {
		imageName = filename[:len(filename)-4]
	}

	if imageName == "folioimg" {
		fmt.Println("Image Found!")
		repo.Image_URL = file["download_url"].(string)
	}
}

// Gets github repos of user
func GetRepos(user string, pass string) ([]Repo, error) {

	username = user
	password = pass

	val, geterr := getHTTP("https://api.github.com/users/" + user + "/repos")
	if geterr != nil {
		return nil, geterr
	}

	var repoList JSONArray
	jsonerr := json.Unmarshal(val, &repoList)

	if jsonerr != nil {
		var e map[string]string
		jsonerror := json.Unmarshal(val, &e)

		if jsonerror != nil {
			return nil, jsonerr
		}

		newerror := GitError{e["message"]}
		return nil, newerror
	}

	repos := make([]Repo, 0, 10)

	for _, v := range repoList {
		repo := ParseRepo(v)
		go GetContentData(&repo)
		repos = append(repos, repo)
	}

	fmt.Printf("")
	return repos, nil
}
