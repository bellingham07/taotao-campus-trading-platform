package gateway

import (
	"bufio"
	"fmt"
	"gateway/config"
	"gateway/registry"
	"net"
	"strconv"
)

type Conf struct {
	Name      string
	Host      string
	Port      int
	Frequency int64
}

type Server struct {
	Host string
	Port int
}

func MustNewServer(serverRegistry registry.ServerRegistry, conf config.Conf) *Server {
	instance, _ := registry.NewDefaultServiceInstance(
		conf.RegistryConf.Name,
		conf.RegistryConf.Host,
		conf.RegistryConf.Port,
		false, nil, "",
	)

	serverRegistry.Register(instance)
	serverRegistry.GetInstances()

	return &Server{}
}

func (s Server) Start() {
	ls, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))
	if err != nil {
		fmt.Printf("start tcp listener error: %v\n", err.Error())
		return
	}
	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Printf("connect error: %v\n", err.Error())
		}
		go func(conn net.Conn) {
			_, err = bufio.NewWriter(conn).WriteString("ok")
			if err != nil {
				fmt.Printf("write conn error: %v\n", err)
			}
		}(conn)
	}
}

func (s Server) Stop() {

}
