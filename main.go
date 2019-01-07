package main

import (
	"log"
	"net/http"
	"os"
)

var (
	RedirectCode = http.StatusFound
	Target       = os.Getenv("REDIRECT_TARGET")
	NoPath       = len(os.Getenv("REDIRECT_NO_PATH")) > 0
	Port         = os.Getenv("PORT")

	OK               = []byte("OK")
	MethodNotAllowed = []byte("Method Not Allowed")
)

func init() {
	if len(os.Getenv("REDIRECT_PERMANENTLY")) > 0 {
		RedirectCode = http.StatusMovedPermanently
	}
	if len(Port) == 0 {
		Port = "80"
	}
}

func main() {
	http.HandleFunc("/.well-known/health", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(OK)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet || req.Method == http.MethodHead {
			if NoPath {
				http.Redirect(w, req, Target, RedirectCode)
			} else {
				http.Redirect(w, req, Target+req.URL.RequestURI(), RedirectCode)
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(MethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
