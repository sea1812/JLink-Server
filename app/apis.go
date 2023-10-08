package app

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

func ApiStopProcess(r *ghttp.Request) {
	//获取进程名称
	mProcess := r.GetString("processname")
	//获取服务器ID
	mUniid := r.GetString("nodeuniid")
	_, er := doStopProcess(mProcess, mUniid)
	if er == nil {
		r.Response.Write("ok")
	} else {
		r.Response.Write(er.Error())
	}
}

func doStopProcess(AProcess string, AUniId string) (g.Map, error) {
	//查询服务器信息，获取地址和PEM文件
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
			mRes := SendCmd_StopProcess(mCertFilename, fmt.Sprint(mNodeInfo["ipaddress"]), AProcess)
			return mRes, nil
		}
	}
}

func SendCmd_StopProcess(ACertFilename string, AUrl string, AProcess string) g.Map {
	mUser := "test" //用户名
	mPwd := "test"  //密码
	mParams := g.Map{
		"process": AProcess,
		"wait":    true,
	}
	mDataMap := g.Map{"cmd": "StopProcess", "params": mParams} //数据明文
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

func ApiStartProcess(r *ghttp.Request) {
	//获取进程名称
	mProcess := r.GetString("processname")
	//获取服务器ID
	mUniid := r.GetString("nodeuniid")
	_, er := doStartProcess(mProcess, mUniid)
	if er == nil {
		r.Response.Write("ok")
	} else {
		r.Response.Write(er.Error())
	}
}

func doStartProcess(AProcess string, AUniId string) (g.Map, error) {
	//查询服务器信息，获取地址和PEM文件
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
			mRes := SendCmd_StartProcess(mCertFilename, fmt.Sprint(mNodeInfo["ipaddress"]), AProcess)
			return mRes, nil
		}
	}
}

func SendCmd_StartProcess(ACertFilename string, AUrl string, AProcess string) g.Map {
	mUser := "test" //用户名
	mPwd := "test"  //密码
	mParams := g.Map{
		"process": AProcess,
		"wait":    true,
	}
	mDataMap := g.Map{"cmd": "StartProcess", "params": mParams} //数据明文
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
