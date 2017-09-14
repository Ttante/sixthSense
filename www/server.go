package www
	
import (
	"fmt"
	"net/http"
)

type Server struct {
	http.Server

}

func New(host string, port int) *Server {
	return &Server{
		Server: http.Server{Addr: fmt.Sprintf("%s:%d", host, port)},
	}
}

func(s *Server) SetHandler(handlers map[string]http.Handler) error {
	mux := http.NewServeMux()
	for k,v := range handlers{
		mux.Handle(k,v)
	}
	s.Handler = mux

	return nil
}
