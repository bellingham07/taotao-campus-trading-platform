package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/service/atcl/cmd/rpc/atclservice"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/file/cmd/api/internal/config"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	UserRpc  userservice.UserService
	CmdtyRpc cmdtyservice.CmdtyService
	AtclRpc  atclservice.AtclService

	Xorm       *xorm.Engine
	FileAtcl   *xorm.Session
	FileCmdty  *xorm.Session
	FileAvatar *xorm.Session

	Oss *Oss
}

type Oss struct {
	Client     *oss.Client
	BaseUrl    string
	BucketName string
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())

	}

	endPoint := c.Oss.EndPoint
	accessKeyId := c.Oss.AccessKeyId
	accessKeySecret := c.Oss.AccessKeySecret
	bucketName := c.Oss.BucketName
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic("[OSS ERROR] NewServiceContext 获取OSS连接错误" + err.Error())
	}

	return &ServiceContext{
		Config:     c,
		UserRpc:    userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		CmdtyRpc:   cmdtyservice.NewCmdtyService(zrpc.MustNewClient(c.CmdtyRpc)),
		AtclRpc:    atclservice.NewAtclService(zrpc.MustNewClient(c.AtclRpc)),
		Xorm:       engine,
		FileAtcl:   engine.Table("file_atcl"),
		FileCmdty:  engine.Table("file_cmdty"),
		FileAvatar: engine.Table("file_avatar"),
		Oss: &Oss{
			Client:     client,
			BaseUrl:    "https://" + bucketName + "." + endPoint + "/",
			BucketName: bucketName,
		},
	}
}
