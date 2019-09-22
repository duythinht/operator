package release

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"go.5kpbs.io/operator/pkg/operator"
	"go.5kpbs.io/operator/pkg/repository"
)

var (
	Command *cobra.Command
	rev     string
	addr    string
	token   string
	submit  bool
)

func init() {
	Command = &cobra.Command{
		Use: "release",
	}
	Command.AddCommand(gitlabCommand())
	Command.PersistentFlags().StringVarP(&rev, "rev", "r", "", "Revision to compare log")
	Command.PersistentFlags().StringVarP(&addr, "addr", "a", "https://git.5kbps.io", "Server address")
	Command.PersistentFlags().StringVarP(&token, "token", "t", "", "Personal Access Token")
	Command.PersistentFlags().BoolVarP(&submit, "submit", "s", false, "Submit a new release")
}

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

func parseRev(rev string) (head string, base string) {
	switch {
	case rev == "":
		return "master", ""
	case strings.Contains(rev, "..."):
		_range := strings.Split(rev, "...")
		return _range[0], _range[1]
	default:
		return rev, ""
	}
}

func exit1(message string) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
	os.Exit(1)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func createGitlabClient(addr string, token string) (*gitlab.Client, error) {

	client := gitlab.NewOAuthClient(nil, token)
	err := client.SetBaseURL(addr)
	if err != nil {
		return nil, err
	}
	return client, nil
}
