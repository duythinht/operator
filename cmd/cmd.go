package cmd

import (
	"github.com/spf13/cobra"
	"go.5kpbs.io/operator/cmd/release"
)

var command *cobra.Command

func init() {
	command = &cobra.Command{Use: "operator"}
	command.AddCommand(release.Command)
}

// Exec run ci-operator command
func Exec() {
	command.Execute()
}
