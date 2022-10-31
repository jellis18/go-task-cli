package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task or tasks as complete",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			}

			err = db.DeleteTask(id)
			fmt.Println("Deleted task", id)
			if err != nil {
				log.Fatal(err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
