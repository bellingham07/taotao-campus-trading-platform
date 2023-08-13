package server

type HttpServer struct {
	Addr         string
	Weight       int
	CWeight      int
	Status       int
	FailWeight   int
	FailCount    int
	RecoverCount int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{
		Addr:    host,
		Weight:  weight,
		CWeight: 0,
	}
}

type HttpServers []*HttpServer

func (s HttpServers) Len() int {
	return len(s)
}

func (s HttpServers) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

func (s HttpServers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
