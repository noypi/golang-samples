package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

func DownloadHandler(res http.ResponseWriter, req *http.Request) {

	var (
		status int
		err    error
	)

	defer func() {
		if nil != err {
			http.Error(res, err.Error(), status)
		}
	}()

	// r.HandleFunc("/download/{name:[a-z]+[0-9]+}/{ext}", DownloadHandler)

	vars := mux.Vars(req)
	name := vars["name"]
	ext := vars["ext"]

	// use path.Base(), might prevent directory traversal attack
	fpath := "./static/fileserver/" + path.Base(name+"."+ext)

	if err = servefile(res, fpath); nil != err {
		status = http.StatusInternalServerError
		return
	}

}
