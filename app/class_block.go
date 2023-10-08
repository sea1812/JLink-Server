package app

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

// TContentBlock 页面内容块对象
type TContentBlock struct {
	BlockTemplate string         //块模版文件
	Data          g.Map          //传入模版数据
	CacheConfig   TCacheConfig   //缓存设置
	CacheIt       bool           //是否启用缓存
	R             *ghttp.Request //浏览器对象
}

// Html 获取页面内容块解析后的文本信息
func (receiver *TContentBlock) Html() string {
	var result string = ""
	if receiver.CacheIt == true {
		//如果启用缓存，则优先从缓存取
		mC, e1 := g.Redis().DoVar("GET", receiver.CacheConfig.CacheName)
		mFc := mC.String() //转换为字符串
		if e1 == nil {
			//取出成功
			if strings.TrimSpace(mFc) == "" {
				//已缓存值为空，则重新解析
				result = receiver.parse()
				//重新置入缓存
				_, _ = g.Redis().DoVar("SET", receiver.CacheConfig.CacheName, result)
				//设置缓存过期时间
				_, _ = g.Redis().DoVar("EXPIRE", receiver.CacheConfig.CacheName, receiver.CacheConfig.GetDuration())
			} else {
				result = mFc
			}
		} else {
			//取出失败，则重新生成页面
			result = receiver.parse()
			//重新置入缓存
			_ = g.DB().GetCache().Set(receiver.CacheConfig.CacheName, result, receiver.CacheConfig.GetDuration())
		}
	} else {
		//如果未启用缓存，则解析页面
		result = receiver.parse()
	}
	return result
}

// parse 从模版生成页面文本
func (receiver *TContentBlock) parse() string {
	result, e1 := receiver.R.Response.ParseTpl(receiver.BlockTemplate, receiver.Data)
	if e1 == nil {
		return result
	} else {
		return ""
	}
}
