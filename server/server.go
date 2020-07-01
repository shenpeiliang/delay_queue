package server

import (
	"context"
	"gin/driver"
	"gin/helper"
	"gin/package/setting"
	"gin/package/validator/extend"
	"gin/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 配置并启动服务
func Run() {
	// gin 运行时 release debug test
	gin.SetMode(setting.ConfigParam.Server.Env)

	//engine := gin.New()
	engine := gin.Default()

	//自定义渲染分隔符
	leftDelims := "{{"
	rightDelims := "}}"
	if setting.ConfigParam.Server.LeftDelims != "" {
		leftDelims = setting.ConfigParam.Server.LeftDelims
	}
	if setting.ConfigParam.Server.RightDelims != "" {
		rightDelims = setting.ConfigParam.Server.RightDelims
	}
	engine.Delims(leftDelims, rightDelims)

	//模板变量使用自定义方法
	engine.SetFuncMap(helper.FunctionHelper)

	// 配置视图
	if "" != setting.ConfigParam.Server.ViewsPattern {
		engine.LoadHTMLGlob(setting.ConfigParam.Server.ViewsPattern)
	}

	//加载静态资源，例如网页的css、js
	engine.Static("/static", "./static")
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	engine.StaticFS("/upload", http.Dir("upload"))
	//加载单个静态文件
	engine.StaticFile("/favicon.ico", "./static/favicon.ico")

	//使用session中间件
	//基于cookie
	//store := cookie.NewStore([]byte(setting.ConfigParam.Session.KeyPairs))

	//基于redis
	//store, err := redis.NewStore(setting.ConfigParam.Redis.DbMaxIdleCconns, "tcp", setting.ConfigParam.Redis.Host+":"+setting.ConfigParam.Redis.Port, setting.ConfigParam.Redis.DbPwd, []byte(setting.ConfigParam.Session.KeyPairs))

	//基于redis连接池
	store, err := redis.NewStoreWithPool(driver.RedisPool, []byte(setting.ConfigParam.Session.KeyPairs))

	if err != nil {
		log.Fatalf(err.Error())
	}
	engine.Use(sessions.Sessions(setting.ConfigParam.Session.Name, store))

	// 注册路由
	router.RegisterRoutes(engine)

	//注册自定义验证器
	extend.New().RegisterValidator()

	serverAddr := setting.ConfigParam.Server.Host + ":" + setting.ConfigParam.Server.Port
	//启动服务
	//err := engine.Run(serverAddr)

	server := &http.Server{
		Addr:           serverAddr,
		Handler:        engine,
		ReadTimeout:    time.Duration(setting.ConfigParam.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.ConfigParam.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: setting.ConfigParam.Server.MaxHeaderBytes,
	}

	//启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Println("start http server and listen ", serverAddr)
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
