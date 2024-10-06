package main

import (
	"flag"
	"fmt"
	"log/slog"
)

var (
	configPath = flag.String(
		"config", "prometheus.yaml",
		"a full path that contains Prometheus YAML file and its name.",
	)
	port = flag.Int(
		"port", 8001,
		"a port number to expose APIs to the network.",
	)
)

func main() {
	flag.Parse()

	config, err := NewConfig(*configPath)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	server, err := NewServer(*port, config)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	slog.Info(fmt.Sprintf("server starting at :%d", *port))
	err = server.Run()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	slog.Info("server terminated")
}
