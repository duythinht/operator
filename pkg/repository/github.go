package repository

import "github.com/google/go-github/github"

type GithubRepository struct {
	Owner               string
	Repository          string
	PersonalAccessToken string
	GithubClient        *github.Client
}
