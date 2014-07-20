package main

import (
	"net/http"
)

func main() {

	// accept POST request only
	http.HandleFunc("/upload", UploadHandler)

	// serve a file
	http.HandleFunc("/favicon.ico", ServeFileHandler)
	http.HandleFunc("/robots.txt", ServeFileHandler)
	http.HandleFunc("/other/path/robots.txt", ServeFileHandler)

	// redirect
	http.Handle("/some/404/path/redirect/to/robots.txt", http.RedirectHandler("/robots.txt", http.StatusMovedPermanently))

	// handle home
	http.HandleFunc("/", HomeHandler) // unknown paths should display this root

	// serve static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/files/"))))

	http.ListenAndServe(":8083", nil)
}
