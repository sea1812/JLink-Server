package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)
import "github.com/gogf/gf/net/ghttp"

type TFrontPage struct {
	BaseTemplate    string       //基本模版
	ContentTemplate string       //内容模版
	FooterTemplate  string       //底部模版
	Data            g.Map        //传入模版数据
	PageTitle       string       //页面标题
	CacheFirst      bool         //是否优先从缓存中加载
	CacheName       string       //缓存名称
	CacheConfig     TCacheConfig //缓存设置
	NeedLogin       bool         //是否需要登录
	LoginUrl        string       //登录页面地址
	NeedPermission  bool         //是否需要检查权限
	PermissionKey   string       //本页面访问权限的KEY值，访问权限在user_info（存储在session中）中的permissons（MAP类型）中
	PermissionUrl   string       //权限不足时跳转的页面
}

//检查是否登录
func (this *TFrontPage) IsLogined(r *ghttp.Request) bool {
	if r.Session.Contains("login_user_info") == true {
		return true
	} else {
		return false
	}
}

//检查是否具有权限
func (this *TFrontPage) IsGranted(r *ghttp.Request) bool {
	if r.Session.Contains("login_user_info") == true {
		//获取permissions字段
		mInfo := r.Session.GetMap("login_user_info")
		mPermission := mInfo["permissions"]
		mv, mValue := mapKeyExist(mPermission.(g.Map), this.PermissionKey)
		if mv == true {
			//判断mValue是否为True
			if mValue == true {
				//已经授权
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

// Display 显示页面
func (this *TFrontPage) Display(r *ghttp.Request) {
	//_ = this.CacheConfig.Load(this.CacheName)
	if this.NeedLogin == true {
		//本页面需要登录，才能查看，检查是否登录
		if this.IsLogined(r) == false {
			//未登录则跳转到loginUrl页面
			r.Response.RedirectTo(this.LoginUrl)
		}
	}
	if this.NeedPermission == true {
		//本页面需要检查权限才能查看，检查权限
		if this.IsGranted(r) == false {
			//未授权则跳转到PermissionUrl
			r.Response.RedirectTo(this.PermissionUrl)
		}
	}
	if this.CacheFirst == true {
		//优先使用缓存，则从缓存中读取页面，读取出错则生成页面并放置到缓存
		mTmp, er := g.DB().GetCache().Get(this.CacheName)
		fmt.Println(mTmp, er)
		if (mTmp == nil) || (er != nil) {
			//读取出错
			fmt.Println("读取出错")
			mPage, er2 := r.Response.ParseTpl(this.BaseTemplate, g.Map{
				"contentTpl": this.ContentTemplate,
				"footerTpl":  this.FooterTemplate,
				"data":       this.Data,
				"pageTitle":  this.PageTitle,
			})
			fmt.Println("mPage=", mPage, er2)
			//如果解析正确，则将mPage放回缓存，并输出mPage
			if er2 == nil {
				_ = g.DB().GetCache().Set(this.CacheName, mPage, this.CacheConfig.GetDuration())
			}
			r.Response.Write(mPage)
		} else {
			//读取正确，输出缓存
			r.Response.Write(mTmp)
		}
	} else {
		//不使用缓存，则直接输出页面
		_ = r.Response.WriteTpl(this.BaseTemplate, g.Map{
			"contentTpl": this.ContentTemplate,
			"footerTpl":  this.FooterTemplate,
			"data":       this.Data,
			"pageTitle":  this.PageTitle,
		})
	}
}
