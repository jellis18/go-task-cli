package cmd

import (
	"fmt"
	"log"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			log.Fatal(err)
		}

		if len(tasks) == 0 {
			fmt.Println("No Tasks to list")
			return
		}

		for _, task := range tasks {
			fmt.Printf("%d: %s\n", task.Key, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
