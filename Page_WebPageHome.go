package main

import (
	"fmt"
	"net/http"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>home</h1>")
		//fmt.Fprint(w, r.URL.Path)
	}
	if r.URL.Path == "/fb" {
		fmt.Fprint(w, "<h1>FB</h1>")
		//fmt.Fprint(w, r.URL.Path)
	}
	if r.URL.Path == "/t" {
		fmt.Fprint(w, "<h1>T</h1>")
		//fmt.Fprint(w, r.URL.Path)
	}

}

func aboutusFunc(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "text/plan")
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

//func fbFunc(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "<h1>FB link</h1>")
//}
//func tFunc(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprint(w, "<h1>T link</h1>")
//}
func pathFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func main() {

	http.HandleFunc("/", homeFunc)
	http.HandleFunc("/aboutus", aboutusFunc)
	http.HandleFunc("/login", loginFunc)
	http.HandleFunc("/logout", logoutFunc)
	http.HandleFunc("/data", dataFunc)
	//http.HandleFunc("/fb", fbFunc)
	//http.HandleFunc("/t", tFunc)
	http.HandleFunc("/path", pathFunc)
	http.ListenAndServe(":3000", nil)
}
