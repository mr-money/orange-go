package App

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"orange-go/Library/Logger"
)

var appLogger = Logger.MustModuleLogger("app")

// PingServer
// @Description: pings the http server to make sure the router is working.
// @param port
// @return error
func PingServer(port string, srv *http.Server) {
	if listen := srv.ListenAndServe(); listen == http.ErrServerClosed {
		return
	}

	for i := 0; i < 5; i++ {
		resp, getErr := http.Get("http://127.0.0.1:" + port + "/")
		if getErr == nil && resp.StatusCode == 200 {
			return
		}

		// Sleep for a second to continue the next ping.
		appLogger.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	appLogger.Panic("Web server cannot connect", "port", port)

}

// Shutdown
// @Description: 优雅关闭服务
// @param srv
func Shutdown(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatal("listen error", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	appLogger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		appLogger.Fatal("Server Shutdown error", "error", err)
	}

	select {
	case <-ctx.Done():
		appLogger.Info("timeout of 3 seconds.")
	}

	appLogger.Info("Server exiting")
}
