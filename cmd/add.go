package cmd

import (
	"bufio"
	"fmt"
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

		add.Add(userInput[:len(userInput)-1], userInput2[:len(userInput2)-1], userInput3[:len(userInput3)-1])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
