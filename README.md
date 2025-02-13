# orange-go

## 简介

> 基于gin开发的微服务web框架，使用ddd领域驱动设计思想为基础架构思想  
> 入口文件可配置启动多个微服务容器  
> 支持Mysql，MongoDB，Redis，Rabbitmq  
> 架构分层为 Model->Repository->Service->App  
> 值对象（Value Object）→ 实体（Entity）→ 领域服务（Domain Service） 

### 初始化框架

#### Golang

> ##### ≥1.17
> 下载地址 https://golang.google.cn/dl/

#### 安装

```git
git clone https://github.com/mr-money/orange-go.git
```

##### 本地：

```shell
go env -w GO111MODULE=on  
go env -w GOPROXY=https://goproxy.cn,direct  
go mod tidy
cd Container/Api
go build -o app
```

##### Docker:
```shell
# --build-arg image={微服务名称}
docker build -t orange-go/api --build-arg image=Api .
# windows宿主机地址：host.docker.internal
docker run -dp 8080:8080 --name orange_go  orange-go/api
```

### 微服务api入口

> 默认API服务入口文件 /Container/Api/main.go   
``` golang
Database.InitMigrate() //数据库迁移
Routes.Include() //初始化路由
```

### 路由

> 默认web路由文件 /Routes/web.go  
> 默认api路由文件路径 /Routes/api.go 需要token中间件验证  
> 支持路由组 中间件 路由规则等方式

### 配置文件

> 默认配置文件 /Config/toml/web.toml.default
>- 同级目录复制 /Config/toml/web.toml.default文件，并修改后缀名为.toml  
   同级目录新建web.go 并写入与配置文件结构相同的struct  
   同级目录config.go中init方法内初始化配置的struct  
   默认数据库配置[DB]、默认Redis配置[REDIS]

### 数据库

数据库连接使用连接池 基于gorm模块连接

#### 数据库连接池配置：

> Library/Gorm/gorm.go:connectMysql()

model层默认demo Model/user.go  
初始化model方法

#### 初始化model方法示例：

``` golang
func (tableName)UserModel() *gorm.DB {
	return Gorm.Mysql.Table(tableName)
}
```

#### repo层使用方法示例：

``` golang
 Model.UserModel().Take(&userInfo, 1)
```

#### 数据迁移

服务启动时运行
> Database.InitMigrate()

ORM加入数据迁移列表方法
> Database.getMysqlMigrations()

#### ORM加入数据迁移列表 例如

``` golang
func getMysqlMigrations() []map[string]interface{} {
	return append(mysqlMigrations,
		//mysql下 user 用户表
		map[string]interface{}{
			"engine":  Model.User{}.GetOption("engine"),
			"comment": Model.User{}.GetOption("comment"),
			"charset": Model.User{}.GetOption("charset"),
			"model":   Model.User{},
		},

		//其他表...
		/*map[string]interface{}{
			"engine":  Model.modelName{}.GetOption("engine"),
			"comment": Model.modelName{}.GetOption("comment"),
			"charset": Model.modelName{}.GetOption("charset"),
			"model":   Model.modelName{},
		},*/
	)
}
```

也可在包Database下新建文件append引入

``` golang
func init() {
	mysqlMigrations := getMysqlMigrations()
	migrations = append(migrations, mysqlMigrations...)
}
```

### 缓存

> Redis 连接基于go-redis模块

#### Redis连接配置

> Library/Cache/redis.go:connectRedis()

#### 生成缓存key

``` golang
Cache.SetKey({key1},{key2},{key3},...)  
```

> 生成缓存key为 key1:key2:key3

#### 使用Redis

``` golang
//获取缓存
Cache.Redis.Get(Cache.Cxt, {key}).Result()

//设置缓存 Cache.Redis后跟Redis原生方法名就可
例如：
  string
  Cache.Redis.Set(
      Cache.Cxt,
      {key}, {data},
      {expires time},
)

```

#### Middle中间件 包MiddleWare：

``` golang
CSRF 防跨站请求伪造
MiddleWare.CSRF() //验证csrf
MiddleWare.CSRFToken() //生成csrf token

MiddleWare.Auth() //jwt登录验证
  auth := context.Request.Header.Get("Authorization")

```

### 队列

> 集成Redis和RabbitMQ  
> 可配置多条队列server监听
> 单条队列任务可配置多个worker消费
 
#### 队列配置
Queue/config.go:initConf() 方法 初始化队列配置
``` golang
return &config.Config{
		DefaultQueue: "orange_go", //默认队列名
		//redis队列
		/*Broker: fmt.Sprintf("redis://%s:%s/%s",
			Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
			Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
			"1",
		),
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},*/

		//rabbitMq队列中间件
		Broker: fmt.Sprintf("amqp://%s:%s@%s:%s",
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "User"),
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Pwd"),
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Host"),
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Port"),
		),
		AMQP: &config.AMQPConfig{
			Exchange:      "orange_go",
			ExchangeType:  "direct",
			BindingKey:    "orange_go_task",
			PrefetchCount: 3,
		},
		ResultBackend: fmt.Sprintf("redis://%s:%s/%s",
			Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
			Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
			"1",
		),
		ResultsExpireIn: 3600, //结果过期时间
	}
```

#### 任务消费配置
Queue/tasks.go:getQueues() 方法 配置队列及相关消费方法
``` golang
// 定义队列消费方法名称
const (
	PrintNameFunc  = "print_name"
	PrintName2Func = "print_name2"
)
type queueGroups struct {
	queueName string                 //队列名称
	tasks     map[string]interface{} //队列下任务 任务名称：任务消费方法
}

//获取队列组配置
func getQueues() *[]queueGroups {

	return &[]queueGroups{
		{
			"queue_test",
			map[string]interface{}{
				PrintNameFunc: QueueDemo.PrintName,
			},
		},
		{
			"queue_test2",
			map[string]interface{}{
				PrintName2Func: QueueDemo.PrintName2,
			},
		},
	}

}
```


#### 加入队列任务

``` golang
//任务参数 数据类型 map
queueParams := make(map[string]interface{})
queueParams["name"] = name

//加入队列任务
res := Queue.AddTask(
		Queue.PrintNameFunc, //消费方法
		queueParams, //任务参数
	)
```

#### 队列任务消费
> 队列消费方法目录路径 Queue|Worker|{对应APP下路径} 

```golang
//name 参数接收 加入任务传入参数
func PrintName(name string) (string, error) {
    //TODO 消费逻辑
	
	if false { //error 3秒重试
		return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
	}
	return name, nil
}

```


