package release

import (
	"github.com/spf13/cobra"
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
	Command.AddCommand(gitlabCommand(), githubCommand())
	Command.PersistentFlags().StringVarP(&rev, "rev", "r", "", "Revision to compare log")
	Command.PersistentFlags().StringVarP(&addr, "addr", "a", "https://git.5kbps.io", "Server address")
	Command.PersistentFlags().StringVarP(&token, "token", "t", "", "Personal Access Token")
	Command.PersistentFlags().BoolVarP(&submit, "submit", "s", false, "Submit a new release")
}
