package www

import (
	"encoding/json"
	"net/http"

	"github.com/ttante/sixthSense/cache"
)

type stats map[string]interface{}

type StatsHandler struct {
	stats []*stats
	dataCache *cache.Cache
}

func NewStats () *StatsHandler {
	return &StatsHandler{
		dataCache: cache.New(),
	}
}

// first () for what class func addresses, 2nd for func params
func(s *StatsHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.get(resp, req)
	// case "POST":
	// 	s.post(resp, req)
	// case "PUT":
	// 	s.put(resp, req)
	// case "DELETE":
	// 	s.delete(resp, req)
	default:
		s.badMethod(resp, req)
	}
}

func(s *StatsHandler) badMethod(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusMethodNotAllowed)
}

func(s *StatsHandler) get(resp http.ResponseWriter, req *http.Request) {

	s.stats = append(s.stats, s.dataCache.Read())

	err := json.NewEncoder(resp).Encode(s.stats)

	if err != nil{
		s.handle500(resp)
	}
}

func(s *StatsHandler) handle500(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusInternalServerError)
}