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
	GroupRoot.ALL("/merch/user", app.MerchUser)             //商家信息页面
	GroupRoot.ALL("/merch/user_edit", app.MerchUser_submit) //商家信息页面
	fmt.Println("Run")
	s.Run()
}
