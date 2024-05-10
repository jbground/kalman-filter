package main

import (
	"github.com/jbground/kalman-filter/web"
	"log"
)

func main() {
	r := web.RoutingByGin()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
