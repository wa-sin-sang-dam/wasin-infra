package main

import "net/http"

type APIHandler interface {
	AddTarget(rw http.ResponseWriter, req *http.Request)
	RemoveTarget(rw http.ResponseWriter, req *http.Request)
	ListTargets(rw http.ResponseWriter, req *http.Request)
}

type apiHandler struct {
	config Config
}

func NewAPIHander(config Config) APIHandler {
	return &apiHandler{
		config: config,
	}
}

func (ah *apiHandler) AddTarget(rw http.ResponseWriter, req *http.Request) {

}

func (ah *apiHandler) RemoveTarget(rw http.ResponseWriter, req *http.Request) {

}

func (ah *apiHandler) ListTargets(rw http.ResponseWriter, req *http.Request) {

}
