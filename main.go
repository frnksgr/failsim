package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/frnksgr/failsim/handler"
)

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	return port
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.DefaultHandler)
	router.HandleFunc("/ping", handler.PingHandler)
	router.HandleFunc("/env", handler.EnvHandler)
	router.HandleFunc("/crash", handler.CrashHandler)
	router.HandleFunc("/sleep/{sec:[0-9]+}", handler.SleepHandler)
	router.HandleFunc("/return/{code:[0-9]+}", handler.ReturnHandler)

	port := getPort()
	fmt.Println("Running server on port ", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
