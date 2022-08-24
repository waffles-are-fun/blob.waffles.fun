/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [list of sha256sums]",
	Short: "Want a waffle? Get a waffle!",
	Long:  `Gets a waffle via sha256sum. The waffle must have been previously uploaded. The waffle will be cached to ~/.waffles/`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("at least one sha256 is required")
		}
		for i, arg := range args {
			if len(arg) != 64 {
				return fmt.Errorf("%s is not a valid sha256sum (item %d)", arg, i)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for i, arg := range args {
			fmt.Println(i, arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
