package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

func GetHandler(res http.ResponseWriter, req *http.Request) {

	var (
		status int
		err    error
	)

	defer func() {
		if nil != err {
			http.Error(res, err.Error(), status)
		}
	}()

	// r.HandleFunc("/get/{html}", GetHandler).Methods("GET")

	vars := mux.Vars(req)
	html := vars["html"]

	// use path.Base() and append '.html', might prevent directory traversal attack :)
	fpath := "./static/html/" + path.Base(html+".html")

	res.Header().Set("Content-Type", "text/html")
	if err = servefile(res, fpath); nil != err {
		status = http.StatusInternalServerError
		return
	}

}
