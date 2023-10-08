package app

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"time"
)

type TCacheConfig struct {
	CacheName          string //缓存名称
	CacheDurationType  string //缓存时间类型
	CacheDurationValue int    //缓存时间值
}

// Load 从数据库加载页面缓存设置
func (this *TCacheConfig) Load(ACacheName string) error {
	res, _ := g.DB().Model("sys_cache_config").Where("cache_name=?", ACacheName).Cache(time.Hour*24, "cache_config_"+ACacheName).One()
	if res.IsEmpty() == false {
		this.CacheName = ACacheName
		this.CacheDurationType = res["cache_duration_type"].String()
		this.CacheDurationValue = res["cache_duration_value"].Int()
		return nil
	} else {
		er := errors.New("TCacheConfig Load 返回为空")
		return er
	}
}

//返回计算好的缓存时间值
func (this *TCacheConfig) GetDuration() time.Duration {
	var m time.Duration
	switch this.CacheDurationType {
	case "hour":
		m = time.Hour * time.Duration(this.CacheDurationValue)
	case "second":
		m = time.Second * time.Duration(this.CacheDurationValue)
	case "minute":
		m = time.Minute * time.Duration(this.CacheDurationValue)
	case "day":
		m = time.Hour * 24 * time.Duration(this.CacheDurationValue)
	}
	return m
}
