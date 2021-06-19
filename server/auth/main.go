package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create the logger:%v", err)
	}
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root123@localhost:27017"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}
	key, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name: "auth",
		Addr: ":8081",
		RegisterFunc: func(server *grpc.Server) {
			authpb.RegisterAuthServiceServer(server, &auth.Service{Logger: logger, OpenIDResolver: &wechat.Service{
				AppSecretID: "3ab4a564c9f74068034e0cdb5c25baf6",
				AppID:       "wxd713b7dafa342293",
			},
				Mongo:          dao.NewMongo(mongoClient.Database("coolcar")),
				TokenExpire:    2 * time.Hour,
				TokenGenerator: token.NewJWTTokenGen("coolcar/auth", privateKey),
			})
		},
		Logger: logger,
	})
	if err != nil {
		logger.Fatal("cannot start server", zap.Error(err))
	}
}
