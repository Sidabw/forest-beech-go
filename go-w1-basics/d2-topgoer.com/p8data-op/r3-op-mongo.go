package data_op

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func R3OpMongo() {
	// 毕竟这个context不知道是啥意思,,,
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println("conn fail on ", mongoUri, err)
		return
	}
	defer func() {
		err := connect.Disconnect(ctx)
		if err != nil {
			fmt.Println("Disconnect fail", err)
			return
		}
	}()

	collection := connect.Database("live_room_setting").Collection("room_setting")
	cur, err := collection.Find(ctx, bson.D{})
	defer cur.Close(ctx)

	if cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println("decode res fail", err)
			return
		}
		fmt.Println("result:", result)
	}
}
