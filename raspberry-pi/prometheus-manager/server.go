// handler.go
package main

import (
	"fmt"
	"net"
	"net/http"
)

// Server exposes APIs listed in endpoint.go.
type Server interface {
	Close()
	Run() error
}

type server struct {
	listner    net.Listener
	handler    http.Handler
	apiHandler APIHandler
	port       int
}

// NewServer instantiate an API handler with the passed port.
func NewServer(port int, config Config) (Server, error) {
	server := &server{
		port:       port,
		apiHandler: NewAPIHander(config),
	}

	err := server.createListener(port)
	if err != nil {
		return nil, err
	}

	server.createHandler()

	return server, nil
}

// Run starts the server and listens on the port passed.
func (s *server) Run() error {
	return http.Serve(s.listner, s.handler)
}

func (s *server) Close() {
	if s.listner == nil {
		return
	}

	_ = s.listner.Close()
}

func (s *server) createListener(port int) error {
	strPort := fmt.Sprintf(":%d", port)
	listner, err := net.Listen("tcp", strPort)
	if err != nil {
		return err
	}

	s.listner = listner

	return nil
}

func (s *server) createHandler() {
	handler := http.NewServeMux()
	handler.HandleFunc(
		string(AddTarget),
		s.apiHandler.AddTarget,
	)
	handler.HandleFunc(
		string(RemoveTarget),
		s.apiHandler.RemoveTarget,
	)
	handler.HandleFunc(
		string(ListTargets),
		s.apiHandler.ListTargets,
	)

	s.handler = handler
}
