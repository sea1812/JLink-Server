package app

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"time"
)

// QueryAllNodes 数据库查询所有节点
func QueryAllNodes() gdb.Result {
	res, _ := g.DB().Model("nodes").OrderAsc("reg_id").Cache(time.Second * 60).All()
	return res
}

func SendCmd_GetAllProcessInfo(ACertFilename string, AUrl string) g.Map {
	mUser := "test"                               //用户名
	mPwd := "test"                                //密码
	mDataMap := g.Map{"cmd": "GetAllProcessInfo"} //数据明文
	mDataJson := gjson.New(mDataMap)
	mData := mDataJson.Export()
	mEncrypt := RSA_Encrypt([]byte(mData), ACertFilename) //对数据进行公钥加密
	//向测试API提交数据
	mR, er := g.Client().Post(AUrl+"/api/cmd", g.Map{
		"user": mUser,
		"pwd":  mPwd,
		"data": string(gbase64.Encode(mEncrypt)), //BASE64编码字符串
	})
	if er == nil {
		//返回API返回信息
		mA := mR.ReadAllString()
		fmt.Println("连接客户端错误=", er)
		mJson := gjson.New(mA)
		return mJson.Map()
	} else {
		return nil
	}
}

// QueryProcesses 查询指定服务器的所有进程信息
func QueryProcesses(AUniId string) (g.Map, error) {
	//查询AUniId对应的服务器地址和证书信息
	mNodeInfo := QueryNodeInfo(AUniId)
	if mNodeInfo == nil {
		//节点不存在，返回错误
		return nil, errors.New("节点不存在")
	} else {
		//节点存在
		//检查证书是否存在，证书路径在./cert/{uni_id}/目录下
		mCertFilename := gfile.Join("./cert", AUniId, fmt.Sprint(mNodeInfo["pem_file"]))
		fmt.Println(mCertFilename)
		if gfile.Exists(mCertFilename) == false {
			//证书文件不存在，返回错误
			return nil, errors.New("数字证书文件不存在")
		} else {
			//证书文件存在，组合指令
			mRes := SendCmd_GetAllProcessInfo(mCertFilename, fmt.Sprint(mNodeInfo["ipaddress"]))
			return mRes, nil
		}
	}
}

//根据 uni_id 查询服务器节点信息
func QueryNodeInfo(AUniId string) g.Map {
	res, _ := g.DB().Model("nodes").Where("uni_id=?", AUniId).One()
	if res.IsEmpty() == false {
		return res.Map()
	} else {
		return nil
	}
}
