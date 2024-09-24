package main

import (
	"flag"
	"fmt"
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
	fmt.Printf("config path passed: %s\n", *configPath)

	config, err := NewConfig(*configPath)
	if err != nil {
		panic(err)
	}

	server, err := NewServer(*port, config)
	if err != nil {
		panic(err)
	}

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
