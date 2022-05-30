package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test.com/tt/api/app"
	"test.com/tt/internal/conf"
	"test.com/tt/internal/logger"
	"time"
)

func Run() {
	conf.InitConfig()
	logger.NewZapLog()
	config := conf.GetConfig()
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger())
	handleRecovery := func(c *gin.Context, err interface{}) {
		c.JSON(500, gin.H{
			"message": err.(string),
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
	ct, _, _ := newController(config)
	app.InitRoute(r.Group("/app/v1"), ct)
	r.GET("/recovery", func(_ *gin.Context) {
		panic("this is a panic")
	})

	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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
