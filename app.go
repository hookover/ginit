package main

import (
	"context"
	"fmt"
	"gin_api/config"
	"gin_api/router"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func init() {
	env := envy.Get("GO_ENV", "local")
	envFile := ".env"

	switch strings.ToLower(env) {
	case "prod":
		envFile = ".env.prod"
	case "sta":
		envFile = ".env.sta"
	default:
		envFile = ".env"
	}

	if err := envy.Load(envFile); err != nil {
		panic(err)
	}

	//初始化日志
	config.SetUpLogger()
	//初始化Redis
	config.SetupRedis()
	//初始化数据库
	config.SetupMysql()
}

func main() {
	appEnv := envy.Get("APP_ENV", "release")
	appPort := envy.Get("APP_PORT", "9090")
	appReadTimeout, _ := strconv.Atoi(envy.Get("APP_READ_TIMEOUT", "120"))
	appWriteTimeout, _ := strconv.Atoi(envy.Get("APP_WRITE_TIMEOUT", "120"))

	if strings.ToLower(appEnv) != "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()

	// 性能分析 - 正式环境不要使用！！！
	pprof.Register(engine)

	// 设置路由
	router.SetupRouter(engine)

	server := &http.Server{
		Addr:         ":" + appPort,
		Handler:      engine,
		ReadTimeout:  time.Duration(appReadTimeout) * time.Second,
		WriteTimeout: time.Duration(appWriteTimeout) * time.Second,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-gin-api             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port:" + appPort + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|      http://127.0.0.1:" + appPort + "       |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
