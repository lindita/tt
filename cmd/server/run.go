package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tt.com/tt/api"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/midware"
)

func Run() {
	config := conf.GetConfig()
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger())
	handleRecovery := func(c *gin.Context, err interface{}) {
		c.JSON(200, gin.H{
			"errorCode": 1,
			"msg": "system error: " + err.(string),
		})
	}
	writer := &lumberjack.Logger{
		Filename:   "./logs/error.log",
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}
	r.Use(gin.CustomRecoveryWithWriter(writer, handleRecovery))
	r.Use(midware.HandleException())
	ttApi, _, _ := NewApi(conf.GetConfig())
	api.InitRoute(r, ttApi)
	fmt.Println("http://127.0.0.1"+config.HTTPServerAddr)
	//https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/
	httpServer := &http.Server{
		Addr:    config.HTTPServerAddr,
		Handler: http.TimeoutHandler(r, time.Second*10, "time out"),
		//ReadTimeout的时间计算是从连接被接受(accept)到request body完全被读取(如果你不读取body，那么时间截止到读完header为止)
		ReadTimeout: 5 * time.Second,
		//从request header的读取结束开始，到 response write结束为止 (也就是 ServeHTTP 方法的声明周期)
		//10秒后返回 Empty reply from server
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen: %s\n", err)
		}
	}()
	waitStopServer(httpServer)

}

func waitStopServer(httpServer *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Err:", err)
	}
	log.Println("Server Shutdown Success！！！")
}
