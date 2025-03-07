package Common

import "net/http"

type Route struct {
	URL         string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
