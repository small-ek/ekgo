package serve

import (
	"context"
	"ekgo/api/boot/config"
	"ekgo/api/boot/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	g errgroup.Group
)

var Engine *gin.Engine

//运行服务A
func RunServe() {
	// 强制日志颜色化
	gin.ForceConsoleColor()
	//加载路由
	Router := router.Load()
	//端口
	address := config.Get.System.Address
	//监听启动服务A,这里可以启动多个服务
	srv := &http.Server{
		Addr:              address,           //端口地址
		Handler:           Router,            //路由
		ReadTimeout:       300 * time.Second, //设置秒的读超时
		WriteTimeout:      300 * time.Second, //设置秒的写超时
		ReadHeaderTimeout: 300 * time.Second, //读取头超时
		IdleTimeout:       300 * time.Second, //空闲超时
		MaxHeaderBytes:    1 << 20,           //HTTP请求的头域最大允许长度1 MB
	}
	//绑定ssl证书或者通过nginx绑定
	//srv.ListenAndServeTLS("./service.key", "./service.pem")
	fmt.Println("  App running at:")
	fmt.Println("  -Local: http://" + address)

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("服务连接失败,可能端口冲突" + err.Error())
		}
	}()

	if err := g.Wait(); err != nil {
		log.Println("启动失败,可能端口冲突请修改配置端口" + err.Error())
	}

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间） )
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务器...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("服务器关闭:" + err.Error())
	}
	log.Println("服务器退出")

}
