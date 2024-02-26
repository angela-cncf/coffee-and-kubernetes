package app

import (
	"fmt"

	"hello-k8s/data"
	"hello-k8s/env"
)

// Hello handles the application logic of this microservice
type Hello struct {
	dataStore data.IDataStore
}

// Initialize configures the application
func (app *Hello) Initialize(config *env.Config) {
	app.dataStore = data.CreateProvider(&config.DB)
}

// Destroy performs any necessary cleanup on nested structures
func (app *Hello) Destroy() {
	app.dataStore.Destroy()
}

// GetHelloMessage computes the greeting that needs to be returned to a user
func (app *Hello) GetHelloMessage(source string) string {
	newCount := app.incrementVisitorCount()
	message := fmt.Sprintf("Welcome %v. You are visitor number %v of our K8s microservice!", source, newCount)
	return message
}

func (app *Hello) incrementVisitorCount() int64 {
	return app.dataStore.IncrementVisitorCount()
}
