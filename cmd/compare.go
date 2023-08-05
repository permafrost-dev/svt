/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/spf13/cobra"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare <version1> <version2>",
	Short: "Compare two semantic versions",
	Long:  `Compare two semantic versions`,
	Run: func(cmd *cobra.Command, args []string) {
		v1, _ := semver.Parse(args[0])
		v2, _ := semver.Parse(args[1])
		result := v1.Compare(v2)

		if result == 0 {
			fmt.Println("eq")
		} else if result == 1 {
			fmt.Println("gt")
		} else {
			fmt.Println("lt")
		}
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
