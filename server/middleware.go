package main

import (
    "fmt"
    "strings"
	"net/http"
    "github.com/google/go-github/v47/github"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
            w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Authentication required"))
			return
		}
        if !strings.HasPrefix(strings.ToLower(auth), "token gho_") {
            w.WriteHeader(http.StatusUnauthorized)
            w.Header().Set("Content-Type", "text/plain")
            w.Write([]byte("Authentication required"))
            return
        }

        transport := OauthTokenTransport{
            Token: auth[6:],
            Transport: http.DefaultTransport,
        }

        client := github.NewClient(transport.Client())

        self, _, err := client.Users.Get(r.Context(), "")
        if err != nil {
            if strings.Contains(err.Error(), "401") {
                w.WriteHeader(http.StatusUnauthorized)
                return
            }
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Println(err)
            return
        }
        membership, _, err := client.Teams.GetTeamMembershipBySlug(r.Context(), "TheMuppets", "waffles", *self.Login)
        if err != nil {
            w.WriteHeader(http.StatusForbidden)
            return
        }
        fmt.Printf("%+v\n", membership)


		next.ServeHTTP(w, r)
	})
}
