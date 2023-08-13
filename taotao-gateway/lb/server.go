package lb

type HttpServer struct {
	Host   string
	Weight int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{
		Host:   host,
		Weight: weight,
	}
}
