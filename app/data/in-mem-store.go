package data

type memory struct {
	count int64
}

// getInMemStore creates a data access service backed by process memory
func getInMemStore() IDataStore {
	memStore := new(memory)
	return memStore
}

// Destroy cleans up local resources for
func (memStore *memory) Destroy() {
}

// IncrementVisitorCount will increment the total number of visits
func (memStore *memory) IncrementVisitorCount() int64 {
	memStore.count++
	return memStore.count
}

// GetVisitorCount returns the total number of visits
func (memStore *memory) GetVisitorCount() int64 {
	return memStore.count
}
