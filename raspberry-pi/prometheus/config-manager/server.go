// handler.go
package main

import (
	"net"
	"net/http"
)

// Server exposes APIs listed in endpoint.go.
type Server interface {
	Close()
	Run() error
}

type server struct {
	listner      net.Listener
	handler      http.Handler
	port         int
	receivedData []byte
}

// NewServer instantiate an API handler with the passed port.
func NewServer(port int) (Server, error) {
	handler := &server{port: port}

	err := handler.createListener(port)
	if err != nil {
		return nil, err
	}

	handler.createHandler()
	return handler, nil
}

// Run starts the server and listens on the port passed.
func (s *server) Run() error {
	go func() {
		_ = http.Serve(s.listner, s.handler)
	}()

	return nil
}

func (s *server) Close() {
	if s.listner == nil {
		return
	}

	_ = s.listner.Close()
}

func (s *server) createListener(port int) error {
	listner, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}

	s.listner = listner

	return nil
}

func (s *server) createHandler() {
	apiHandler := NewAPIHander()
	handler := http.NewServeMux()
	handler.HandleFunc(
		string(AddTarget),
		apiHandler.AddTarget,
	)
	handler.HandleFunc(
		string(RemoveTarget),
		apiHandler.RemoveTarget,
	)
	handler.HandleFunc(
		string(ListTargets),
		apiHandler.ListTargets,
	)

	s.handler = handler
}
