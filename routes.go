package main

import (
	"git.rappet.de/rappet/netmap-api/routing"
)

var routes = routing.Routes{
	routing.Route{
		"List all PTRs",
		"GET",
		"/ptrs",
		routing.JsonHttp(PtrIndex),
	},
	routing.Route{
		"Get PTR by IP Address",
			"GET",
			"/ptrs/{address}",
		routing.JsonHttp(GetPtr),
	},
}
