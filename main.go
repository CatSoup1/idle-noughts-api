package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)


func main() {
	url := os.Getenv("url")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	col := client.Database("Cluster1").Collection("idle-nought")

	r := gin.Default()

	r.GET("/get/leaderboard", func(c *gin.Context) {
		cur, err := col.Find(context.TODO(), bson.D{{}})
		if err != nil {
			panic(err)
		}
		var lbs []primitive.M

		for cur.Next(context.TODO()) {
			var lb bson.M
			err := cur.Decode(&lb)
			if err != nil {
				panic(err)
			}
			lbs = append(lbs, lb)
		}
		defer cur.Close(context.TODO())

		c.JSON(200, lb)
	})

	r.Run(":8080")
}
