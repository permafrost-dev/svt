/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/spf13/cobra"
)

// latestCmd represents the latest command
var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Compare two semantic versions and print the latest",
	Long:  `Compare two semantic versions and print the latest`,
	Run: func(cmd *cobra.Command, args []string) {
		v1, _ := semver.Parse(args[0])
		v2, _ := semver.Parse(args[1])
		result := v1.Compare(v2)

		if result == 0 {
			fmt.Println(v1.String())
		} else if result == 1 {
			fmt.Println(v1.String())
		} else {
			fmt.Println(v2.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// latestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// latestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
