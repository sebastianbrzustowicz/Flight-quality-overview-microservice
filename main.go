package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/upload-data", uploadHandler)
	http.ListenAndServe(":8083", nil)
}
