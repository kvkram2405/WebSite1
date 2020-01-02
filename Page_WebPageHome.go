package main

import (
	"fmt"
	"net/http"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home</h1>")
	fmt.Fprint(w, "<B>Kalyan </B>")
}
func aboutusFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>AboutUs... OUR COMPANY</h1>")
}
func loginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Login...to check your data Enter Your login name Password here </h1>")
}
func logoutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>logout</h1>")
}
func dataFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>GetData From Database</h1>")
}
func fbFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>FB link</h1>")
}


func main() {
	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/aboutus", aboutusFunc)
	http.HandleFunc("/login", loginFunc)
	http.HandleFunc("/logout", logoutFunc)
	http.HandleFunc("/data", dataFunc)
	http.HandleFunc("/fb", fbFunc)
	http.ListenAndServe(":3000", nil)
}
