package registry

import (
	"fmt"
	"gateway/config"
	"gateway/server"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

type ConsulRegistry struct {
	ListenOn            string
	client              *api.Client
	localServerInstance ServerInstance
	HttpServerMap       map[string]server.HttpServers
	FetchInterval       int64
}

func (c *ConsulRegistry) Register(serverInstance ServerInstance) {
	c.localServerInstance = serverInstance
	schema, tags := "http", make([]string, 0)

	id := serverInstance.GetKey() + "-" + serverInstance.GetHost() + "-" + strconv.Itoa(serverInstance.GetPort())
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    serverInstance.GetKey(),
		Address: serverInstance.GetHost(),
		Port:    serverInstance.GetPort(),
	}

	if serverInstance.IsSecure() {
		tags = append(tags, "secure=true")
	} else {
		tags = append(tags, "secure=false")
	}
	if serverInstance.GetMetadata() != nil {
		var tags []string
		for key, value := range serverInstance.GetMetadata() {
			tags = append(tags, key+"="+value)
		}
		registration.Tags = tags
	}
	registration.Tags = tags

	// 增加consul健康检查回调函数
	check := &api.AgentServiceCheck{
		TCP:                            c.ListenOn,
		Timeout:                        "10s",
		Interval:                       "30s",
		DeregisterCriticalServiceAfter: "20s", // 故障检查失败20s后 consul自动将注册服务删除
		//HTTP:
	}

	registration.Check = check

	if serverInstance.IsSecure() {
		schema = "https"
	}
	check.HTTP = schema + "://" + registration.Address + ":" + strconv.Itoa(registration.Port) + "/actuator/health"

	// 注册服务到consul
	if err := c.client.Agent().ServiceRegister(registration); err != nil {
		log.Fatalf("[FATAL REGISTRY] 网关注册失败 %v", err)
	}
}

func (c *ConsulRegistry) Deregister() {
	if c.localServerInstance == nil {
		return
	}
	_ = c.client.Agent().ServiceDeregister(c.localServerInstance.GetID())
	c.localServerInstance = nil
}

func NewConsulRegistry(conf *config.Conf) *ConsulRegistry {
	if len(conf.RegistryConf.Host) < 3 {
		log.Fatalf("[FATAL REGISTRY] 网关注册失败 check host\n")
	}

	if conf.RegistryConf.Port <= 0 || conf.RegistryConf.Port > 65535 {
		log.Fatalf("[FATAL REGISTRY] 网关注册失败 check port, port should between 1 and 65535\n")
	}

	apiConfig := api.DefaultConfig()
	apiConfig.Address = conf.RegistryConf.Host + ":" + strconv.Itoa(conf.RegistryConf.Port)
	apiConfig.Token = conf.RegistryConf.Token
	client, err := api.NewClient(apiConfig)
	if err != nil {
		log.Fatalf("[FATAL REGISTRY] 网关注册失败 %v", err)
	}

	listenOn := conf.GatewayConf.Host + strconv.Itoa(conf.GatewayConf.Port)
	return &ConsulRegistry{client: client, ListenOn: listenOn, FetchInterval: conf.RegistryConf.Frequency}
}

func (c *ConsulRegistry) GetInstances() {
	var ticker = time.NewTicker(time.Duration(c.FetchInterval) * time.Second)
	c.HttpServerMap["cmdty.rpc"] = server.HttpServers{}
	for {
		select {
		case <-ticker.C:
			for location, _ := range c.HttpServerMap {
				if err := c.discovery(location); err != nil {

				}
			}
		}
	}
}

func (c *ConsulRegistry) discovery(serviceName string) error {
	//catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	servers, _, err := c.client.Health().Service(serviceName, "", false, nil)
	if err != nil {
		log.Printf("[ERROR DISCOVERY] 获取 %v 服务失败 %v\n", serviceName, err)
	}
	fmt.Println(servers)
	//if len(servers) > 0 {
	//	result := make([]ServerInstance, len(servers))
	//	for index, sever := range servers {
	//		s := DefaultServerInstance{
	//			ID:       sever.ServiceID,
	//			Key:     sever.ServiceName,
	//			Host:     sever.Address,
	//			Port:     sever.ServicePort,
	//			Metadata: sever.ServiceMeta,
	//		}
	//		result[index] = s
	//	}
	//	return nil
	//}
	return nil
}
