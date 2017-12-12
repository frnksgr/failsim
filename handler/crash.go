package handler

import (
  "fmt"
  "net/http"
  "os"
)

func CrashHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Killing me")
  os.Exit(1)
}

