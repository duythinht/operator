package release

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"go.5kpbs.io/operator/pkg/operator"
	"go.5kpbs.io/operator/pkg/repository"
)

func gitlabCommand() *cobra.Command {

	var (
		pid  string
		head string
		base string
	)

	cmd := &cobra.Command{
		Use: "gitlab",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := createGitlabClient(addr, token)
			exitIfError(err)

			operator := &operator.Operator{
				Repo: &repository.GitlabRepository{
					PID:          pid,
					GitlabClient: client,
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
			if len(args) != 1 {
				exit1("You must enter repository to be release!!!\n")
			}
			pid = args[0]
			head, base = parseRev(rev)
		},
	}

	return cmd
}

func createGitlabClient(addr string, token string) (*gitlab.Client, error) {

	client := gitlab.NewOAuthClient(nil, token)
	err := client.SetBaseURL(addr)
	if err != nil {
		return nil, err
	}
	return client, nil
}
