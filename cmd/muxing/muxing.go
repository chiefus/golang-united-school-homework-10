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

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", greetingUrlPartHandler).Methods("GET")
	router.HandleFunc("/bad", serverErrorHandler).Methods("GET")
	router.HandleFunc("/data", receiveParameterHandler).Methods("POST")
	router.HandleFunc("/headers", headersSumHandler).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(defaultHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func greetingUrlPartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s!", vars["PARAM"])
}

func serverErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func receiveParameterHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := r.Body
	body, err := io.ReadAll(requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "I got message:\n")
	w.Write(body)
}

func headersSumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.Header.Get("a")
	bStr := r.Header.Get("b")
	a, err := strconv.Atoi(aStr)
	if err != nil {
		log.Fatalln(err)
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		log.Fatalln(err)
	}
	sum := a + b
	sumStr := strconv.Itoa(sum)
	w.Header().Add("a+b", fmt.Sprint(sumStr))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
