package main

import (
    "net/http"
    "fmt"
    )



func main() {

   http.HandleFunc("/",handler)
   http.ListenAndServe(":8080",nil)

 }
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, this is my first GoLang program! %s",r.URL.Path[1:])
}


