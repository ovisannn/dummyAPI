package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func userCollection(c *mongo.Database) {
	collection = c.Collection("user")
}
