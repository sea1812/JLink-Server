package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	_ "github.com/mattn/go-sqlite3"
	"goJLinkServer/app"
)

func main() {
	app.SysConfig = app.NewTSysConfig()
	//_ = app.SysConfig.LoadFromDB(true)
	//设置Redis缓存，暂不启用
	//adapterRedis := adapter.NewRedis(g.Redis())
	//g.DB().GetCache().SetAdapter(adapterRedis)
	s := g.Server()
	s.AddStaticPath("/static", "./static")
	s.BindStatusHandler(404, app.Page404) //绑定状态页
	//根路径组
	GroupRoot := s.Group("/")
	GroupRoot.ALL("/", app.IndexPage)                       //显示首页
	GroupRoot.ALL("/processes/{uni_id}", app.ProcessPage)   //显示客户端进程信息页面
	GroupRoot.ALL("/api/startprocess", app.ApiStartProcess) //启动进程
	GroupRoot.ALL("/api/stopprocess", app.ApiStopProcess)   //停止进程
	fmt.Println("Run")
	s.Run()
}
