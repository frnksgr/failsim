package handler

import (
//  "fmt"
  "os"
  "strings"
  "net/http"
  "encoding/json"
)

func EnvHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  envs := make(map[string]string)
  for _, env := range os.Environ() {
    kv := strings.Split(env, "=")
    envs[kv[0]] = kv[1]
  }
  data := struct {
    RequestHeader map[string][]string
    RequestMethod string
    RequestProtocol string
    ProcessEnvironment map[string]string
  } {
    r.Header,
    r.Method,
    r.Proto,
    envs,
  }
  json.NewEncoder(w).Encode(data)
}
