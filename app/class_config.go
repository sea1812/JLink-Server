package app

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"time"
)

type TSysConfig struct {
	data g.Map
}

func NewTSysConfig() *TSysConfig {
	return &TSysConfig{
		data: make(g.Map),
	}
}

// LoadFromDB  从数据库加载全部设置参数，参数为是否使用缓存
func (this *TSysConfig) LoadFromDB(Cached bool) error {
	var res gdb.Result
	var er error
	if Cached == true {
		res, er = g.DB().Model("sys_config").Cache(time.Hour*24, sysConfigCacheName).All()
	} else {
		res, er = g.DB().Model("sys_config").All()
	}
	if er == nil {
		for _, v := range res {
			this.data[v["Key"].String()] = v["Value"].String()
		}
	} else {
		this.data = g.Map{}
	}
	return er
}

// ClearCache 清除缓存
func (this *TSysConfig) ClearCache() error {
	_, err := g.DB().GetCache().Remove(sysConfigCacheName)
	return err
}

// Query 通过键值查询全局设置
func (this *TSysConfig) Query(AKey string) (error, string) {
	var exists bool
	var rr interface{}
	var er error
	var ra string
	exists, rr = mapKeyExist(this.data, AKey)
	if exists == false {
		er = errors.New("键值不存在")
		ra = ""
	} else {
		ra = fmt.Sprint(rr)
		er = nil
	}
	return er, ra
}
