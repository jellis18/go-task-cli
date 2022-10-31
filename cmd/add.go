package cmd

import (
	"fmt"
	"log"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
