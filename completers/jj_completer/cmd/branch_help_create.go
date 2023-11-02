package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_help_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_createCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_createCmd)
}