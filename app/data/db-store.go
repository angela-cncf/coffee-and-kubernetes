package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"hello-k8s/env"
)

type postgres struct {
	db              *gorm.DB
	visitorCountKey string
}

// getDBStore creates a data acccess service backed by the configured database
func getDBStore(config *env.DBConfig) IDataStore {
	dbStore := new(postgres)

	if len(config.VisitorCountKey) > 0 {
		dbStore.visitorCountKey = config.VisitorCountKey
	} else {
		dbStore.visitorCountKey = "shared-visitors-key"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("failed to connect to db with info %+v: %+v", psqlInfo, err)
		panic("failed to connect database")
	}

	db.LogMode(true)
	db.AutoMigrate(Visitor{})
	dbStore.db = db

	return dbStore
}

// Destroy is a destructor for DataAccessService
func (dbStore *postgres) Destroy() {
	dbStore.db.Close()
}

// IncrementVisitorCount will increment the total number of visits
func (dbStore *postgres) IncrementVisitorCount() int64 {
	count := dbStore.GetVisitorCount() + 1
	// Try to create the record, just in case it's the very first time anyone has visited this microservice
	if err := dbStore.db.Table("visitors").Where(&Visitor{ID: dbStore.visitorCountKey}).FirstOrCreate(
		&Visitor{
			ID:    dbStore.visitorCountKey,
			Count: count,
		}).Error; err != nil {
		fmt.Printf("Error trying to create/update all visitors count in DB for %v: %v", dbStore.visitorCountKey, err.Error())
	}
	// Update the visitor count
	attributes := map[string]interface{}{"Count": count}
	if err := dbStore.db.Model(&Visitor{}).Where(&Visitor{ID: dbStore.visitorCountKey}).Updates(attributes).Error; err != nil {
		fmt.Printf("Error trying to increment all visitors count in DB: %v", err.Error())
	}
	return count
}

// GetVisitorCount returns the total number of visits
func (dbStore *postgres) GetVisitorCount() int64 {
	var allVisitors Visitor
	if err := dbStore.db.Table("visitors").Where(&Visitor{ID: dbStore.visitorCountKey}).First(&allVisitors).Error; err != nil {
		fmt.Printf("Error trying to get all visitors count in DB: %v", err.Error())
		return 0
	}
	return allVisitors.Count
}
