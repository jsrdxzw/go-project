package dao

import "go.mongodb.org/mongo-driver/mongo"

type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{col: db.Collection("trip")}
}

type TripRecord struct {
}
