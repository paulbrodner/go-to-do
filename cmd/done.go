/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/paulbrodner/go-to-do/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Closing a task",
	Long:  `Mark a task as done`,
	Run:   doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(dataFile)

	id, err2 := strconv.Atoi(args[0])
	if err2 != nil {
		log.Fatalln(args[0], "is not a valid label\n", err2)
	}
	if id > 0 && id < len(items) {
		items[id-1].Done = true
		fmt.Printf("%q %v\n", items[id-1].Text, "marked as done")
		todo.SaveItems(dataFile, items)
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
