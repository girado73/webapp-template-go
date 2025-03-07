package main

import (
	"github.com/gorilla/mux"
	"webapp-template-go/src/Testmodule"
)

var routes = Testmodule.RouteCollector

func GetRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		r.HandleFunc(route.URL, route.HandlerFunc).Methods(route.Method)
	}

	return r
}
