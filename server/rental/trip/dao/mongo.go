package dao

import (
	"coolcar/shared/mgutil"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField      = "trip"
	accountIDField = tripField + ".accountid"
	statusField    = tripField + ".status"
)

type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{col: db.Collection("trip")}
}

type TripRecord struct {
	mgutil.IDField
	mgutil.UpdatedAtField
	Trip
}
