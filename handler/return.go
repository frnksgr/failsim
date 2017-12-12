package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ReturnHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	code, err := strconv.Atoi(mux.Vars(r)["code"])
	if err != nil {
		fmt.Println("wrong format")
	}
	if code > 0 {
		fmt.Println("Returning http response code", code)
		w.WriteHeader(code)
	}
}
