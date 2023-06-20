package Log

import (
	"context"
	"go-study/Model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logModel Model.Log

//
// GetLogPage
// @Description: 分页获取ccgp爬虫数据
// @param jobId
// @return []Model.Ccgp
//
func GetLogPage(page int64, pageSize int64) []Model.Log {
	filter := primitive.M{
		"type": 1,
	}

	//分页
	findOptions := options.Find().
		SetLimit(pageSize).
		SetSkip((page - 1) * pageSize)

	var logData []Model.Log

	cur, _ := logModel.Collection().Find(
		context.Background(),
		filter,
		findOptions,
	)

	_ = cur.All(context.Background(), &logData)

	if len(logData) == 0 {
		logData = []Model.Log{}
	}

	return logData
}

//
// CountLog
// @Description: 获取log数量
// @param jobId
// @return int64
//
func CountLog() int64 {

	filter := primitive.M{
		"type": 1,
	}

	count, err := logModel.Collection().CountDocuments(context.Background(), filter)

	if err != nil {
		return 0
	}

	return count
}
