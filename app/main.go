package main

import (
	"flag"
	"fmt"

	"hello-k8s/api"
	"hello-k8s/env"
)

func main() {
	fmt.Println("Starting up Hello K8s")

	// Get the microservice's configuration
	configFile := flag.String("config", "", "the config file")
	flag.Parse()

	fmt.Println("Reading the configuration from " + *configFile)
	serviceConfig := env.CreateConfig(*configFile)

	// Create a REST API server
	rest := new(api.RESTServer)
	defer rest.Destroy()

	// Configure and start the REST API
	fmt.Println(fmt.Sprintf("Initializing microservice using %+v", serviceConfig))
	rest.Initialize(serviceConfig)
	rest.Run()
}
