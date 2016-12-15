package getgit

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
)

func NilString(n interface{}) string {
	if n == nil {
		return ""
	} else {
		return n.(string)
	}
}

func ParseRepo(v map[string]interface{}) Repo {
	id := int(v["id"].(float64))
	name := v["name"].(string)
	full_name := NilString(v["full_name"])
	url := NilString(v["html_url"])
	description := NilString(v["description"])
	created, _ := time.Parse(time.RFC3339,v["created_at"].(string))
	updated, _ := time.Parse(time.RFC3339,v["updated_at"].(string))
	homepage := NilString(v["homepage"])
	stars := int(v["stargazers_count"].(float64))
	watchers := int(v["watchers_count"].(float64))
	language := NilString(v["language"])
	return Repo{
		id,
		name,
		full_name,
		url,
		description,
		created,
		updated,
		homepage,
		stars,
		watchers,
		language,
	}
}

func GetRepos(user string) ([]Repo, error) {
	
	res, geterr := http.Get("https://api.github.com/users/" + user + "/repos")
	if geterr != nil {
		return nil, geterr
	}
	
	val, readerr := ioutil.ReadAll(res.Body)
	if readerr != nil {
		return nil, readerr
	}

	var repo []map[string]interface{}
	jsonerr := json.Unmarshal(val, &repo)
	
	if jsonerr != nil {
		var e map[string]string
		jsonerror := json.Unmarshal(val, &e)
		
		if jsonerror != nil {
			return nil, jsonerr				
		}
		
		newerror := GitError{e["message"]}
		return nil, newerror
	}

	repos := make([]Repo,0,10)

	for _, v := range repo {
		repos = append(repos, ParseRepo(v))
	}

	fmt.Printf("")
	return repos, nil
}
