package main

import (
	"GoCode/myapp"
	"net/http"
)
func main() {

    http.ListenAndServe(":3000", myapp.NewHttpHandler())
}