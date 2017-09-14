package www

import (
	"encoding/json"
	"net/http"
)

type StatsHandler struct {
	stats []* stats
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

func(s *StatsHandler) get(resp http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(resp).Encode(s.stats)

	if err != nil{
		s.handle500(resp)
	}
}

func(s *StatsHandler) handle500(resp http.ResponseWriter) {
	resp.WriteHeader(http.InternalServerError)
}