package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, err := url.Parse("http://xxxxx.com:8080")
	if err != nil {
		log.Println("bad url:", err)
		return
	}
	httpProxy := httputil.NewSingleHostReverseProxy(targetURL)
	http.Handle("/", httpProxy)
	http.ListenAndServe(":8080", nil)
	log.Println("program started")
}
