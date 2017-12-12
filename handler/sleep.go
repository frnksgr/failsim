package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func SleepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	sec, err := strconv.Atoi(mux.Vars(r)["sec"])
	if err != nil {
		fmt.Println("wrong format")
	}
	if sec > 0 {
		fmt.Println("Sleeping for", sec, "seconds")
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
