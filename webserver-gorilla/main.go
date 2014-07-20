package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// try to filter files to download, example only
	r.HandleFunc("/download/{name:[a-z]+[0-9]+}/{ext}", DownloadHandler)

	// set content type to html example
	r.HandleFunc("/get/{html}", GetHandler).Methods("GET")

	//-
	http.Handle("/", r)

	// wait for clients
	http.ListenAndServe(":8083", nil)
}
