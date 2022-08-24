package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

func main() {

	logger := httplog.NewLogger("main", httplog.Options{
		JSON: true,
	})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	r.Use(middleware.Heartbeat("/ping"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/get/{digest}", func(r chi.Router) {
		r.Get("/", getByDigest)
	})

	r.Route("/upload", func(r chi.Router) {
		//TODO: auth middleware
		r.Post("/", uploadByDigest)
	})

	http.ListenAndServe(":5000", r)
}

func getByDigest(w http.ResponseWriter, r *http.Request) {
	digest := chi.URLParam(r, "digest")
	if len(digest) != 64 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("https://blob.waffles.fun/%s/%s", digest[0:2], digest), 302)
}

func uploadByDigest(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	buf, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	digest := fmt.Sprintf("%x", sha256.Sum256(buf))

	os.Mkdir(fmt.Sprintf("%s/%s", "/tmp", digest[0:2]), 0755)
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s", "/tmp", digest[0:2], digest), buf, 0644)
	if err != nil {
		panic(err)
	}
}
