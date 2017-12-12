package util

import (
  "net/http"
)

func RequestTarget(r *http.Request) string {
  target := r.Header.Get("X-Forwarded-Proto")
  if target == "" {
    target = r.URL.Scheme
    if target == "" {
      target = "http" // default
    }
  }
  return target + "://" + r.Host
}
