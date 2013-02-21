package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func routerHandlerFunc(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world")
}

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	return router
}

func main() {
	handler := routerHandlerFunc(router())
    err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
