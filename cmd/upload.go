/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:    "upload [list of files]",
	Short:  "waffles a waffle to the waffle",
	PreRun: NeedAuth,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("at least one file is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		config, err := LoadConfig()
		if err != nil {
			panic(err)
		}

		manifest, err := LoadManifest()
		if err != nil {
			panic(err)
		}

		client := &http.Client{}

		for _, file := range args {
			dat, err := os.ReadFile(file)

			if err != nil {
				panic(err)
			}

			req, err := http.NewRequest("POST", "http://localhost:5000/upload", bytes.NewBuffer(dat))
			req.Header.Add("Authorization", "Token "+config.Token)

			resp, err := client.Do(req)

			if err != nil {
				panic(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != 200 {
				panic(fmt.Sprintf("failed to upload file, got status code %d with body %s", resp.StatusCode, body))
			}

			fmt.Printf("uploaded %s: status %d result %s", file, resp.StatusCode, body)

			// append to manifest.txt
			manifest[file] = string(body)
			manifest.Save()

		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
