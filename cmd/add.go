package cmd

import (
	"fmt"
	"bufio"
	"os"

	add "github.com/hahwul/hack-pet/pkg/adding"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add snippet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("[command]")
		fmt.Print(">>> ")
		userInput, _ := reader.ReadString('\n')
		fmt.Println(userInput)

		fmt.Println("[desc]")
		fmt.Print(">>> ")
		userInput2, _ := reader.ReadString('\n')
		fmt.Println(userInput2)

		fmt.Println("[toml filename | e.g nmap_full_scan.toml]")
		fmt.Print(">>> ")
		userInput3, _ := reader.ReadString('\n')
		fmt.Println(userInput3)

		Add(userInput,userInput2,userInput3)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
