package main

import (
	"fmt"
	"net/http"
	"github.com/ttante/sixthSense/www"
)

func main() {
	fmt.Println("started")
	srv := www.New("0.0.0.0", 8000)

	routes := map[string]http.Handler{
		"/stats": new(www.StatsHandler),
	}

	srv.SetHandler(routes)

	srv.ListenAndServe()
}