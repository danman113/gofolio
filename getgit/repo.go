package getgit

import (
	"time"
)

type Repo struct {
	id          int       `json:"id"`
	name        string    `json:"name"`
	full_name   string    `json:"full_name"`
	url         string    `json:"html_url"`
	description string    `json:"description"`
	created     time.Time `json:"created_at"`
	updated     time.Time `json:"updated_at"`
	homepage    string    `json:"homepage"`
	stars       int       `json:"stargazers_count"`
	watchers    int       `json:"watchers_count"`
	language    string    `json:"language"`
}

type GitError struct {
	Message string
}

func (err GitError) Error() string {
	return err.Message
}
