// endpoint.go lists out all accessible API endpoints and its URl
package main

type Endpoint string

const (
	AddTarget    Endpoint = "POST /api/target/add"
	RemoveTarget Endpoint = "POST /api/target/remove"
	ListTargets  Endpoint = "GET /api/target/list"
)
