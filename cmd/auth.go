/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
    "net/http"
	"github.com/cli/oauth/device"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Refresh/delete authentication",
	Long: `waffles requires a github auth token. we walk you through generating one here.
note: this token _will be_ sent to hte server during uploads. We request "org:read" permissions to
verify which teams you are a member of and if you have the ability to waffle.

We store a sha256 of your token that includes which teams you were in when the token was initalized.
If teams change you may need to refresh your token.
`,
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
