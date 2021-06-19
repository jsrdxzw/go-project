package dao

import (
	"context"
	mongo3 "coolcar/shared/mgutil"
	"coolcar/shared/mgutil/testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
	"time"
)

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "60bb7aac7a4054ab769f5531",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "60bb7aac7a4054ab769f5532",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "60bb7aac7a4054ab769f5571",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Errorf("cannot connect mongodb:%v", err)
	}
	m := NewMongo(connect.Database("coolcar"))

	mock(t, m, ctx)

	// new case
	m.newObjID = func() primitive.ObjectID {
		objID, _ := primitive.ObjectIDFromHex("60bb7aac7a4054ab769f5571")
		return objID
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("failed resolve account id for %q:%v", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve account id err want: %q, got: %q", cc.want, id)
			}
		})
	}
}

func mock(t *testing.T, m *Mongo, ctx context.Context) {
	_, err := m.col.InsertMany(ctx, []interface{}{
		bson.M{
			mongo3.IDField: mustObjID("60bb7aac7a4054ab769f5531"),
			openIDField:    "openid_1",
		},
		bson.M{
			mongo3.IDField: mustObjID("60bb7aac7a4054ab769f5532"),
			openIDField:    "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert values: %v", err)
	}
}

func mustObjID(hex string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return objectID
}

// TestMain
// It is sometimes necessary for a test program to do extra setup or teardown before or after testing
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
