package svc

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
	"go-go-zero/service/atcl/cmd/rpc/atclservice"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/file/cmd/api/internal/config"
	"go-go-zero/service/file/cmd/api/internal/middleware"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm       *xorm.Engine
	FileAtcl   *xorm.Session
	FileCmdty  *xorm.Session
	FileAvatar *xorm.Session

	UserRpc  userservice.UserService
	CmdtyRpc cmdtyservice.CmdtyService
	AtclRpc  atclservice.AtclService

	Oss *Oss

	JwtAuth rest.Middleware
}

type Oss struct {
	Client     *oss.Client
	BaseUrl    string
	BucketName string
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.FileApi.Mysql)

	endPoint := c.FileApi.Oss.EndPoint
	accessKeyId := c.FileApi.Oss.AccessKeyId
	accessKeySecret := c.FileApi.Oss.AccessKeySecret
	bucketName := c.FileApi.Oss.BucketName
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic("[OSS ERROR] NewServiceContext 获取OSS连接错误" + err.Error())
	}

	fmt.Println("sadasdasdasdfoisfgvsdjklvhdvkldfhv")

	return &ServiceContext{
		Config:     c,
		Xorm:       engine,
		FileAtcl:   engine.Table("file_atcl"),
		FileCmdty:  engine.Table("file_cmdty"),
		FileAvatar: engine.Table("file_avatar"),
		UserRpc:    userservice.NewUserService(zrpc.MustNewClient(c.FileApi.UserRpc)),
		CmdtyRpc:   cmdtyservice.NewCmdtyService(zrpc.MustNewClient(c.FileApi.CmdtyRpc)),
		AtclRpc:    atclservice.NewAtclService(zrpc.MustNewClient(c.FileApi.AtclRpc)),
		JwtAuth:    middleware.NewJwtAuthMiddleware().Handle,
		Oss: &Oss{
			Client:     client,
			BaseUrl:    "https://" + bucketName + "." + endPoint + "/",
			BucketName: bucketName,
		},
	}
}
