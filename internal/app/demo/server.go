package demo

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"service-template-go/internal/app/demo/controller"
)

func init() {
	logWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/gin.log",
		MaxSize:    200, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	// todo change mode
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.LoggerWithWriter(logWriteSyncer))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rgV1 := r.Group("/v1")
	controller.AddUserRoutes(rgV1)

	// todo change port
	err := r.Run()
	if err != nil {
		return
	}
}
