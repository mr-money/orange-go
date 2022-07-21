# go-study

## 简介
> 基于gin开发的微服务web框架，使用ddd领域驱动设计思想为基础架构思想  
> 架构分层为 Model->Repository->Service->App  
> 值对象（Value Object）→ 实体（Entity）→ 领域服务（Domain Service）  
> 数据库连接池基于gorm开发  

### 初始化框架
#### go版本
> ##### 1.17 https://golang.google.cn/dl/

项目根目录运行：
> go env -w GO111MODULE=on  
> go env -w GOPROXY=https://goproxy.cn,direct  
> go mod tidy


### 入口
> 默认入口文件 /main.go   
> Routes.Include()方法初始化路由

### 路由
> 默认web路由文件 /Routes/web.go  
> todo 默认api路由文件路径 /Routes/api.go 需要token中间件验证  
> 支持路由组 中间件 路由规则等方式

### 配置文件
> 默认配置文件 /Config/toml/web.toml.default  
>- 同级目录复制 /Config/toml/web.toml.default文件，并修改后缀名为.toml  
   同级目录新建web.go 并写入与配置文件结构相同的struct  
   同级目录config.go中init方法内初始化配置的struct

### 数据库
数据库连接使用连接池 基于gorm连接

> 数据库连接池配置：  
> Library/Gorm/gorm.go:connectMysql()  
>
> 数据库配置按照上文配置文件方法配置后  
> model层默认demo Model/user.go  
> 初始化model方法 
#### 初始化model方法示例：
```
func UserModel() *gorm.DB {
	return Gorm.Mysql.Table(tableName)
}
```
#### 数据迁移
Database/migration.go:migration()


#### repo层使用方法示例：
```
 Model.UserModel().Take(&userInfo, 1)
```

#### Middle中间件 包MiddleWare：
```
CSRF 防跨站请求伪造
MiddleWare.CSRF() //验证csrf
MiddleWare.CSRFToken() //生成csrf token

MiddleWare.Auth() //jwt登录验证
  auth := context.Request.Header.Get("Authorization")

```



