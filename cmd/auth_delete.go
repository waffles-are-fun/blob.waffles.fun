/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete your waffle token, you will need to auth again to waffle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	authCmd.AddCommand(deleteCmd)
}
