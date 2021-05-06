package main

import (
  "log"

  "github.com/vijayviji/proglog/internal/server"
)

func main() {
  srv := server.NewHTTPServer("0.0.0.0:8080")
  log.Println("Listening on Port 8080")
  log.Fatal(srv.ListenAndServe())
}