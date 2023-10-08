package app

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

// 商户档案管理
func MerchUser(r *ghttp.Request) {
	//获取商户信息
	mMerchUsers, _ := g.DB().Model("pa_merch_users").All()
	mDateTime := time.Now().Format("20160906")
	a := TFrontPage{}
	a.PageTitle = "商户档案管理"
	a.CacheName = "merch_user"               //缓存名字
	a.CacheFirst = false                     //不使用缓存
	a.CacheConfig.CacheDurationValue = 24    //缓存时间
	a.CacheConfig.CacheDurationType = "hour" //缓存时间单位
	a.CacheConfig.CacheName = "merch_user"   //设置项中记录的缓存名字（暂时没用，保留）
	a.BaseTemplate = "admin_base.html"       //基类模版名称
	a.ContentTemplate = "/merch/user.html"   //内容模版名称
	a.FooterTemplate = "index.html"          //脚模版名称
	a.Data = g.Map{
		"activeband": "商户管理", //注入到页面模版的变量
		"activeitem": "商户信息设置",
		"list":       mMerchUsers,
		"filename":   mDateTime,
	}
	a.NeedLogin = false      //不需要登录
	a.NeedPermission = false //不需要检查权限
	a.Display(r)
}

// 提交商户档案
func MerchUser_submit(r *ghttp.Request) {
	a := TFrontPage{}
	a.PageTitle = "商户档案管理"
	a.CacheName = "merch_user"                  //缓存名字
	a.CacheFirst = false                        //不使用缓存
	a.CacheConfig.CacheDurationValue = 24       //缓存时间
	a.CacheConfig.CacheDurationType = "hour"    //缓存时间单位
	a.CacheConfig.CacheName = "merch_user"      //设置项中记录的缓存名字（暂时没用，保留）
	a.BaseTemplate = "admin_base.html"          //基类模版名称
	a.ContentTemplate = "/merch/user_edit.html" //内容模版名称
	a.FooterTemplate = "index.html"             //脚模版名称
	a.Data = g.Map{
		"activeband": "商户管理", //注入到页面模版的变量
		"activeitem": "商户信息设置",
	}
	a.NeedLogin = false      //不需要登录
	a.NeedPermission = false //不需要检查权限
	a.Display(r)
}
