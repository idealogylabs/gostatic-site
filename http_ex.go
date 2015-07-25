package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.StripPrefix("/", http.FileServer(http.Dir("assets/")))
  http.Handle("/", fs)

  log.Println("Listening at port 3000")
  http.ListenAndServe(":3000", nil)
}