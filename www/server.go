package www
	
import (
	"net/http"
)

type Server struct {
	http.Server

	Host string
	Port int

}

func New(host string, port int) *Server {
	return &Server{
		Host:host,
		Port:port,
	}
}

func(s *Server) SetHandler(handlers map[string]http.Handler) error {
	mux := http.NewServeMux
	for k,v := range handlers{
		mux.Handle(k,v)
	}
	s.Handler = mux
}
