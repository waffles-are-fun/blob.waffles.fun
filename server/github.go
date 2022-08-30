package main

import (
    "net/http"
)

type OauthTokenTransport struct {
	Token     string
	Transport http.RoundTripper
}

func (t *OauthTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Token "+t.Token)
	return t.Transport.RoundTrip(req)
}

func (t *OauthTokenTransport) Client() *http.Client {
    return &http.Client{Transport: t}
}
