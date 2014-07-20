package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func UploadHandler(res http.ResponseWriter, req *http.Request) {

	var (
		status int
		err    error
	)

	defer func() {
		if nil != err {
			http.Error(res, err.Error(), status)
		}
	}()

	// parse request
	const _24K = (1 << 20) * 24
	if err = req.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		return
	}

	for _, fheaders := range req.MultipartForm.File {
		for _, hdr := range fheaders {

			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				return
			}

			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./uploaded/" + hdr.Filename); nil != err {
				status = http.StatusInternalServerError
				return
			}

			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}

			res.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))

		}
	}

}
