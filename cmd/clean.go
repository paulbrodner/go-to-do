/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean all TODO tasks",
	Long:  `Delete the local configuration file with all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove(dataFile)
		fmt.Println("All tasks are deleted!")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
