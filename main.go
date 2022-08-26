package main

import (
	"context"
	"fmt"
	"go-study/Database"
	"go-study/Routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//
//  main
//  @Description: 入口
//
func main() {

	//默认服务
	defaultServer()

}

//
// defaultServer
// @Description: 默认服务
//
func defaultServer() {
	//数据库迁移
	Database.InitMigrate()

	// 加载路由
	Routes.Include(
		Routes.Web, //默认web路由
		Routes.Api, //api路由，需要token中间件验证
	)

	port := "8080"

	//启动服务
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Routes.GinEngine,
	}

	//优雅关闭
	shutdown(srv)

	//启动自检
	pingServer(port, srv)
}

//
// pingServer
// @Description: pings the http server to make sure the router is working.
// @param port
// @return error
//
func pingServer(port string, srv *http.Server) {
	if listen := srv.ListenAndServe(); listen == http.ErrServerClosed {
		return
	}

	for i := 0; i < 5; i++ {
		resp, getErr := http.Get(fmt.Sprintf("http://127.0.0.1:%s/", port))
		if getErr == nil && resp.StatusCode == 200 {
			return
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	log.Panicln("Web server in port " + port + ", Cannot connect to this server!")

}

//
// shutdown
// @Description: 优雅关闭服务
// @param srv
//
func shutdown(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}

	log.Println("Server exiting")
}
