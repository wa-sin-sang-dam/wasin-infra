// endpoint.go lists out all accessible API endpoints and its URl
package main

type Endpoint string

const (
	AddTarget    Endpoint = "/api/target/add"
	RemoveTarget Endpoint = "/api/target/remove"
	ListTargets  Endpoint = "/api/target/list"
)
