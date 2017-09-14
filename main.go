package main

import (
	"fmt"

	"github.com/ttante/sixthSense/www"
)

func main() {
	fmt.Println("started")
	srv := www.New("0.0.0.0", 8000)
}