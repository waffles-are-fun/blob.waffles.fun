/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "waffles all the waffles from the waffle",
	Run: func(cmd *cobra.Command, args []string) {
		manifest, err := LoadManifest()
		if err != nil {
			panic(err)
		}

		config, err := LoadConfig()
		if err != nil {
			panic(err)
		}

		os.MkdirAll(config.CacheDir, 0755)

		for path, digest := range manifest {
			cachedPath := fmt.Sprintf("%s/%s", config.CacheDir, digest)

			if _, err := os.Stat(cachedPath); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("%s not found, downloading", digest)
			}

			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("hardlinking %s to %s", digest, path)
				err := os.Link(cachedPath, path)
				if err != nil {
					panic(err)
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
