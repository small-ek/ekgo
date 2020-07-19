package serve

import (
	"context"
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

type Server struct {
	Address           string
	ReadTimeout       int64
	WriteTimeout      int64
	ReadHeaderTimeout int64
	IdleTimeout       int64
	MaxHeaderBytes    int
	Router            *gin.Engine
}

//运行服务A
func (this *Server) RunServe() {
	// 强制日志颜色化
	gin.ForceConsoleColor()
	//监听启动服务A,这里可以启动多个服务
	srv := &http.Server{
		Addr:              this.Address,                                        //端口地址
		Handler:           this.Router,                                         //路由
		ReadTimeout:       time.Duration(this.ReadTimeout) * time.Second,       //设置秒的读超时
		WriteTimeout:      time.Duration(this.WriteTimeout) * time.Second,      //设置秒的写超时
		ReadHeaderTimeout: time.Duration(this.ReadHeaderTimeout) * time.Second, //读取头超时
		IdleTimeout:       time.Duration(this.IdleTimeout) * time.Second,       //空闲超时
		MaxHeaderBytes:    this.MaxHeaderBytes,                                 //HTTP请求的头域最大允许长度1 MB
	}
	//绑定ssl证书或者通过nginx绑定
	//srv.ListenAndServeTLS("./service.key", "./service.pem")
	fmt.Println("  App running at:")
	fmt.Println("  -Local: http://" + this.Address)

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
