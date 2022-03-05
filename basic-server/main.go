package main

import (
	"net/http"
	"os"

	"github.com/richardsnider/go-examples/basic-server/util"
)

func main() {
	util.Log("Server starting . . .")
	serverError := http.ListenAndServe(":"+os.Args[1], http.HandlerFunc(util.HttpHandler))
	util.Log("Listening on port " + os.Args[1] + " . . .")

	if serverError != nil {
		util.Log(serverError)
		os.Exit(1)
	}
}
