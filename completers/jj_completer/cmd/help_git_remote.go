package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_remoteCmd).Standalone()

	help_gitCmd.AddCommand(help_git_remoteCmd)
}