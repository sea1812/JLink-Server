package app

import "github.com/gogf/gf/frame/g"

//系统设置缓存名称
const sysConfigCacheName = "cache_sys_config"

// SysConfig 系统设置全局变量
var SysConfig *TSysConfig

//检查MAP中键值是否存在
func mapKeyExist(mapValue g.Map, key string) (bool, interface{}) {
	if value, ok := mapValue[key]; ok {
		return true, value
	} else {
		return false, nil
	}
}
