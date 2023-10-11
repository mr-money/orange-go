package MongoDB

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v1/log"
	"go-study/Config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

var (
	MongoDataBase = connect()
	client        *mongo.Client
)

//
// Connect
// @Description: mongo连接
// @return *mongo.Database
//
func connect() *mongo.Database {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接uri
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		//"mongodb://%s:%s",
		Config.GetFieldByName(Config.Configs.Web.MongoDB, "User"),
		Config.GetFieldByName(Config.Configs.Web.MongoDB, "Pwd"),
		Config.GetFieldByName(Config.Configs.Web.MongoDB, "Host"),
		Config.GetFieldByName(Config.Configs.Web.MongoDB, "Port"),
	)

	// 构建mongo连接可选属性配置
	opt := new(options.ClientOptions)
	// 设置最大连接的数量
	opt = opt.SetMaxPoolSize(10)
	// 设置连接超时时间 5000 毫秒
	du, _ := time.ParseDuration("5000")
	opt = opt.SetConnectTimeout(du)
	// 设置连接的空闲时间 毫秒
	mt, _ := time.ParseDuration("5000")
	opt = opt.SetMaxConnIdleTime(mt)

	// 开启驱动
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), opt)
	if err != nil {
		log.INFO.Panicln(err)
	}

	// 注意，在这一步才开始正式连接mongo
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.INFO.Panicln(err)
	}

	log.INFO.Println("MongoDB Database [" + Config.GetFieldByName(Config.Configs.Web.MongoDB, "Port") + "]: Connect Success!")

	return client.Database(Config.GetFieldByName(Config.Configs.Web.MongoDB, "DbName"))
}

// Transaction
// @Description: 事务操作
// @param fun
// @return interface{}
// @return error
func Transaction(fun func()) {
	//开启事务
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)
	session, _ := client.StartSession()
	defer session.EndSession(context.TODO())

	_, _ = session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		fun()
		return nil, nil
	}, txnOptions)

	return
}
