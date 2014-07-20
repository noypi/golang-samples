package main

import (
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/html/home.html")
}
