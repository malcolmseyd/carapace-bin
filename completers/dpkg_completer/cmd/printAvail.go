package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var printAvailCmd = &cobra.Command{
	Use:   "print-avail",
	Short: "Display available version details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printAvailCmd).Standalone()

	carapace.Gen(printAvailCmd).PositionalAnyCompletion(
		apt.ActionPackages(),
	)
}
