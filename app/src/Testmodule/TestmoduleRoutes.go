package Testmodule

import (
	"webapp-template-go/src/Common"
)

const RESOURCES = "/testmodule"
const RESOURCE = "/testmodule/{id}"

var RouteCollector = Common.Routes{
	Common.Route{URL: RESOURCES, Method: "GET", HandlerFunc: getAllTestmodules},
	Common.Route{URL: RESOURCES, Method: "POST", HandlerFunc: createSingleTestmodule},
	Common.Route{URL: RESOURCE, Method: "GET", HandlerFunc: getOneTestmodule},
}
