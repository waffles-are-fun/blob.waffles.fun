/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var authRefreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "waffle a new token to waffle with",
	Run: func(cmd *cobra.Command, args []string) {
		err := RefreshAuth()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	authCmd.AddCommand(authRefreshCmd)
}
