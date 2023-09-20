package App

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//
// PingServer
// @Description: pings the http server to make sure the router is working.
// @param port
// @return error
//
func PingServer(port string, srv *http.Server) {
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
// Shutdown
// @Description: 优雅关闭服务
// @param srv
//
func Shutdown(srv *http.Server) {
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
