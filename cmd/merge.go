package cmd

import (
	"fmt"

	merge "github.com/hahwul/hack-pet/pkg/merging"
	"github.com/spf13/cobra"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge all snippets (/snippets/* => hack-pet.toml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("merge called")
		if len(args) >= 1 {
			merge.Merge(args[0])
		} else {
			merge.Merge("snippets")
		}
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
