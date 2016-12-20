package getgit

import (
	"time"
)

type Repo struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Full_name    string    `json:"full_name"`
	URL          string    `json:"html_url"`
	Description  string    `json:"description"`
	Created      time.Time `json:"created_at"`
	Updated      time.Time `json:"updated_at"`
	Homepage     string    `json:"homepage"`
	Stars        int       `json:"stargazers_count"`
	Watchers     int       `json:"watchers_count"`
	Language     string    `json:"language"`
	Git_URL      string    `json:"git_url"`
	Contents_URL string    `json:"contents_url"`
	Image_URL    string    `json:"image"`
}

type GitError struct {
	Message string
}

func (err GitError) Error() string {
	return err.Message
}
