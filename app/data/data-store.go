package data

import (
	"fmt"

	"hello-k8s/env"
)

// IDataStore defines the interface required to access data storage
type IDataStore interface {
	Destroy()

	IncrementVisitorCount() int64
	GetVisitorCount() int64
}

// CreateProvider will create and configure a data storage provider, as configured in the environment settings
func CreateProvider(dbConfig *env.DBConfig) IDataStore {
	if dbConfig.Enabled == true {
		fmt.Println("Connecting to DB")
		return getDBStore(dbConfig)
	}
	fmt.Println("Storing data in memory")
	return getInMemStore()
}
