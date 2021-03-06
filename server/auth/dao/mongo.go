package dao

import (
	"context"
	"coolcar/shared/mgutil"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

type Mongo struct {
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

// NewMongo creates a new mongo dao
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("account"),
		newObjID: primitive.NewObjectID,
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, openId string) (string, error) {
	res := m.col.FindOneAndUpdate(
		c,
		bson.M{
			openIDField: openId,
		},
		mgutil.SetOnInsert(bson.M{
			mgutil.IDField: m.newObjID(),
			openIDField:    openId,
		}),
		options.FindOneAndUpdate().SetUpsert(true),
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate:%v", err)
	}
	var row mgutil.IDField
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result:%v", err)
	}
	return row.ID.Hex(), nil
}
