package main

import (
	"git.rappet.de/rappet/netmap-api/routing"
)

var routes = routing.Routes{
	routing.Route{
		"List all PTRs",
		"GET",
		"/ptrs",
		PtrIndex,
	},
}
