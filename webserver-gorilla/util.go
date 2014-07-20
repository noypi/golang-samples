package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func servefile(res http.ResponseWriter, fpath string) (err error) {
	outfile, err := os.OpenFile(fpath, os.O_RDONLY, 0x0444)
	if nil != err {
		return
	}

	// 32k buffer copy
	written, err := io.Copy(res, outfile)
	if nil != err {
		return
	}

	log.Println("served file:", outfile.Name(), ";length:", written)
	return
}
