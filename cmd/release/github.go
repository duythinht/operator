package release

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"go.5kpbs.io/operator/pkg/operator"
	"go.5kpbs.io/operator/pkg/repository"
	"golang.org/x/oauth2"
)

func githubCommand() *cobra.Command {

	var (
		owner string
		repo  string
		head  string
		base  string
	)

	cmd := &cobra.Command{
		Use: "github",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := createGithubClient(addr, token)
			exitIfError(err)

			operator := &operator.Operator{
				Repo: &repository.GithubRepository{
					Owner:        owner,
					Repository:   repo,
					GithubClient: client,
				},
			}

			if !submit {
				cl, err := operator.Changelog(head, base)
				exitIfError(err)
				fmt.Println(cl)
				return
			}
			exitIfError(operator.Release(head, base))
			fmt.Printf("A release was submitted!!!\n")
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			head, base = parseRev(rev)
		},
	}
	cmd.PersistentFlags().StringVarP(&owner, "owner", "o", "duythinht", "Repository owner")
	cmd.PersistentFlags().StringVarP(&repo, "repository", "x", "", "Repository")
	return cmd
}

func createGithubClient(addr string, token string) (*github.Client, error) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(context.Background(), ts)

	// Make base API url should be correct for github.com and selfhost
	if addr != "https://git.5kbps.io" && addr != "https://github.com" && addr != "http://github.com" {
		baseURL := fmt.Sprintf("%s/api/v3/", addr)
		return github.NewEnterpriseClient(baseURL, baseURL, tc)
	}
	return github.NewClient(tc), nil
}
