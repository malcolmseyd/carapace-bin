package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var multi_pack_indexCmd = &cobra.Command{
	Use:     "multi-pack-index",
	Short:   "Write and verify multi-pack-indexes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(multi_pack_indexCmd).Standalone()
	multi_pack_indexCmd.Flags().String("batch-size", "", "during repack, collect pack-files of smaller size into a batch that is larger than this size")
	multi_pack_indexCmd.Flags().String("object-dir", "", "object directory containing set of packfile and pack-index pairs")
	multi_pack_indexCmd.Flags().Bool("progress", false, "force progress reporting")
	rootCmd.AddCommand(multi_pack_indexCmd)
}
