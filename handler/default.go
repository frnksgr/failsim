package handler

import (
	"fmt"
	"net/http"

	"github.com/frnksgr/failsim/util"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "High There")

	urls := []string{
		"/",
		"/ping",
		"/env",
		"/crash",
		"/sleep/<seconds>",
		"/return/<http-response-code>",
	}

	target := util.RequestTarget(r)

	for _, url := range urls {
		fmt.Fprintln(w, target+url)
	}
}
