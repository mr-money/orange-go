package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"orange-go/Library/MongoDB"
)

/*use orange_go

db.log.insertOne({
    "title":"mongo test",
    "content":"aaaaaaaaaaaaaa",
    "type":1,
    "created_at":"2023-06-20 16:16:48",
})*/

// Log
// @Description: 日志表
type Log struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Content   string             `json:"content" bson:"content"`
	Title     string             `json:"title" bson:"title"`
	Type      int                `json:"type" bson:"type"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}

// Collection
// @Description:创建集合
// @receiver Log
// @return *mongo.Collection
func (Log) Collection() *mongo.Collection {
	return MongoDB.MongoDataBase.Collection("log")
}
