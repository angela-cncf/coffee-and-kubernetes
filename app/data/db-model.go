package data

// Visitor defines the database entry for a visitor
type Visitor struct {
	ID    string `gorm:"primary_key"`
	Count int64
}
