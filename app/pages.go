package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Page404 404状态页
func Page404(r *ghttp.Request) {
	r.Response.WriteTpl("404.html")
}

// ProcessPage 显示服务器详细进程页面
func ProcessPage(r *ghttp.Request) {
	mUniid := r.GetVar("uni_id").String()
	fmt.Println(mUniid)
	a := TFrontPage{}
	a.PageTitle = "详细进程信息"
	a.CacheName = "processes_page_" + mUniid             //缓存名字
	a.CacheFirst = false                                 //不使用缓存
	a.CacheConfig.CacheDurationValue = 24                //缓存时间
	a.CacheConfig.CacheDurationType = "hour"             //缓存时间单位
	a.CacheConfig.CacheName = "processes_page_" + mUniid //设置项中记录的缓存名字（暂时没用，保留）
	a.BaseTemplate = "admin_base.html"                   //基类模版名称
	a.ContentTemplate = "processes.html"                 //内容模版名称
	a.FooterTemplate = "empty.html"                      //脚模版名称
	mInfo, _ := QueryProcesses(mUniid)
	a.Data = g.Map{
		"activeband": "nodes", //注入到页面模版的变量
		"activeitem": "",
		"processes":  mInfo, //获取全部服务器节点列表
		"nodeinfo":   QueryNodeInfo(mUniid),
	}
	a.NeedLogin = false      //不需要登录
	a.NeedPermission = false //不需要检查权限
	a.Display(r)
}

// IndexPage 首页(服务器列表页)
func IndexPage(r *ghttp.Request) {
	a := TFrontPage{}
	a.PageTitle = "服务器集群"
	a.CacheName = "index_page"               //缓存名字
	a.CacheFirst = false                     //不使用缓存
	a.CacheConfig.CacheDurationValue = 24    //缓存时间
	a.CacheConfig.CacheDurationType = "hour" //缓存时间单位
	a.CacheConfig.CacheName = "index_page"   //设置项中记录的缓存名字（暂时没用，保留）
	a.BaseTemplate = "admin_base.html"       //基类模版名称
	a.ContentTemplate = "index.html"         //内容模版名称
	a.FooterTemplate = "index.html"          //脚模版名称
	a.Data = g.Map{
		"activeband": "nodes", //注入到页面模版的变量
		"activeitem": "",
		"nodes":      QueryAllNodes(), //获取全部服务器节点列表
	}
	a.NeedLogin = false      //不需要登录
	a.NeedPermission = false //不需要检查权限
	a.Display(r)
}
