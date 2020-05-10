package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("🎂"));
  })

  log.Println("Starting HTTP Server...")
  log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
