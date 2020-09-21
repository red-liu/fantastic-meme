package main

import (
    "fmt"
    "net/http"
    "my.com/firstmod"
    "my.com/secondmod"
)

func main() {
    http.HandleFunc("/firstmod", firstmodHandler)
    http.HandleFunc("/secondmod", secondmodHandler)
    http.ListenAndServe("0.0.0.0:8080", nil)
}

func firstmodHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, firstmod.HelloV2("world"))
}


func secondmodHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, secondmod.Hello())
}
