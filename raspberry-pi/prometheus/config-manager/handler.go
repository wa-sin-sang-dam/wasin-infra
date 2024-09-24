package main

import "net/http"

type APIHandler interface {
	AddTarget(rw http.ResponseWriter, req *http.Request)
	RemoveTarget(rw http.ResponseWriter, req *http.Request)
	ListTargets(rw http.ResponseWriter, req *http.Request)
}
