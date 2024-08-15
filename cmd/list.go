/*
Copyright © 2024 SHIVAM VIJAYWARGI <vijaywargishivam@gmail.com>
*/
package cmd

import (
	"github.com/shivamvijaywargi/go-todo-term/todos"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the todos",
	Long: `This command will let you list all the todos that you have stored in the todos.csv file.
  For example:

  ./go-todo-term list`,
	Run: func(cmd *cobra.Command, args []string) {
		todos.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
