package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func handleName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["PARAM"]
	response := fmt.Sprintf("Hello, %s!", param)
	_, err := fmt.Fprint(w, response)
	if err != nil {
		panic(err)
	}
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	param, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	response := fmt.Sprintf("I got message:\n%s", param)
	_, err = fmt.Fprint(w, response)
	if err != nil {
		panic(err)
	}
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	//host := "127.0.0.1"
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
