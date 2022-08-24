package main

import (
    "fmt"
    "net/http"
    "github.com/cli/oauth/device"
)

func main() {

    clientID := "c5bacf7d056bf8b7772e"
    scopes := []string{"read:org"}

    code, err := device.RequestCode(http.DefaultClient, "https://github.com/login/device/code", clientID, scopes)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Copy code: %s\n", code.UserCode)
    fmt.Printf("then open: %s\n", code.VerificationURI)

    accessToken, err := device.PollToken(http.DefaultClient, "https://github.com/login/oauth/access_token", clientID, code)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Access token: %s\n", accessToken.Token)
}
