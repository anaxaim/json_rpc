package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func main() {
	srv := rpc.NewServer()
	srv.RegisterCodec(json2.NewCodec(), "application/json")
	srv.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	srv.RegisterService(new(GreetingService), "GreetingService")

	router := mux.NewRouter()
	router.Handle("/", srv)

	log.Fatal(http.ListenAndServe(":8000", router))
}
