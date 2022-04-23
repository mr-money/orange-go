package Index

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"go-study/Config"
	"go-study/Library/Func"
	Go_redis "go-study/Library/Go-redis"
	"go-study/Model"
	"net/http"
)

//
// Home
// @Description: 控制器主页
// @param c
//
func Home(c *gin.Context) {
	c.String(http.StatusOK, "index page")
}

//
// Middle
// @Description: 中间件
// @param c
//
func Middle(c *gin.Context) {
	req := c.Query("request")
	fmt.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

//
// Conf
// @Description: 读取配置
// @param c
//
func Conf(c *gin.Context) {
	var webConfig Config.Web

	_, err := toml.DecodeFile("./Config/web.toml", &webConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("config")
	fmt.Println(webConfig)
	c.JSON(200, gin.H{"config": webConfig})

}

//
// Database
// @Description: 数据库连接池
// @param c
//
func Database(c *gin.Context) {
	var userInfo Model.User
	//userInfo = User.FindById(1)

	c.JSON(200, gin.H{
		"config": Config.Configs.Web,
		"user":   userInfo,
	})
}

//redis连接
func RedisCon(c *gin.Context) {
	/*rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})*/

	redis := Go_redis.Redis

	ctx := context.Background()

	val, _ := redis.Get(ctx, "go_redis").Result()

	val1, _ := redis.HGet(ctx, "dn:trick:user:11111", "msg_send_flag").Result()

	c.JSON(200, gin.H{
		"res":  val,
		"val1": val1,
	})
}

func JsonToStruct(c *gin.Context) {

	str := `{"Name": "Ed", "Text": "Knock knock."}`

	var structRes struct {
		Name string
		Text string
	}

	Func.JsonToStruct(str, &structRes)

	fmt.Printf("structRes的类型是%T", structRes)
	fmt.Println(structRes)

	c.String(http.StatusOK, "JsonToStruct")
}
