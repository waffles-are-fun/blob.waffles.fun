/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/cli/oauth/device"
	"github.com/spf13/cobra"
	"net/http"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "waffling waffles to the waffle requires waffles",
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func NeedAuth(cmd *cobra.Command, args []string) {
	conf, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	if conf.Token == "" {
		fmt.Println("Auth is not present, please authenticate: ")
		err := RefreshAuth()
		if err != nil {
			panic(err)
		}
	}
}

func RefreshAuth() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	clientID := "c5bacf7d056bf8b7772e"
	scopes := []string{"read:org"}

	code, err := device.RequestCode(http.DefaultClient, "https://github.com/login/device/code", clientID, scopes)
	if err != nil {
		return err
	}

	fmt.Printf("Copy code: %s\n", code.UserCode)
	fmt.Printf("then open: %s\n", code.VerificationURI)

	accessToken, err := device.PollToken(http.DefaultClient, "https://github.com/login/oauth/access_token", clientID, code)
	if err != nil {
		return err
	}

	config.Token = accessToken.Token
	err = config.Save()
	if err != nil {
		return err
	}

	fmt.Println("Token saved!")
	return nil
}
