package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func english(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func telugu(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Namaste, %s!\n", ps.ByName("name"))
}
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name/english", english)
	router.GET("/hello/:name/telugu", telugu)
	http.ListenAndServe(":3000", router)
}
