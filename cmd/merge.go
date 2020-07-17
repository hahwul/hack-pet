package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	merge "github.com/hahwul/hack-pet/pkg/merge"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
