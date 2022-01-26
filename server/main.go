package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"restfulAPI/handlers"
)

func main () {
	router()
}

var routes = []struct{
	Method string
	Path string
	handler http.HandlerFunc
}{
	{"GET", 	"/hello", handlers.Hello},
	{"GET", 	"/v1/user-posts/{id}", handlers.UserPosts},
}


func router(){
	r := mux.NewRouter()

	for _ ,routeCfg := range routes {
		route := r.Handle(routeCfg.Path, routeCfg.handler)
		route.Methods(routeCfg.Method)
	}

	fmt.Println("server is running at localhost:1234")
	log.Fatal(http.ListenAndServe(":1234", r))
}