package main

import (
  "fmt"
  "net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}

func main() {
  Start()
}
