# go-study

##简介
> 基于gin开发的微服务web框架，使用ddd领域驱动设计思想为基础架构思想  
> 架构分层为 Repository->Model->Service->App  
> 数据库连接池基于gorm开发  
> 值对象（Value Object）→ 实体（Entity）→ 领域服务（Domain Service）


### 入口
> 默认入口文件 /main.go 
>- Routes.Include()方法初始化路由

### 路由
> 默认web路由文件 /Routes/web.go  
> todo 默认api路由文件路径 /Routes/api.go 需要token中间件验证  
> 支持路由组 中间件 路由规则等方式

### 配置文件
> 默认配置文件 /Config/web.toml.default  
>- 同级目录复制 /Config/web.toml.default文件，并修改后缀名为.toml  
   同级目录新建web.go 并写入与配置文件结构相同的struct  
   同级目录config.go中配置struct

