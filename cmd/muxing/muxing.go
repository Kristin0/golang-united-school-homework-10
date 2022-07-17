package main

import (
	"fmt"
	"io/ioutil"
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

	router.HandleFunc("/name/{PARAM}", ParamGetHandler)
	router.HandleFunc("/bad", GetBad)
	router.HandleFunc("/data", PostData)
	router.HandleFunc("/headers", PostHeaders)


	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}
func PostHeaders(w http.ResponseWriter, r *http.Request){
	//w.WriteHeader(http.StatusOK)
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Set("a+b", strconv.Itoa(a+b))
}

func PostData(w http.ResponseWriter, r *http.Request){

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	
	w.Write([]byte(fmt.Sprint("I got message:\n", string(body))))
}

func GetBad(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(500)
}

func ParamGetHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint("Hello, ", vars["PARAM"], "!")))

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
