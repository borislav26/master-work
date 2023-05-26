package main

import (
	"log"
	"mailer-service/builder"
)

func main() {
	services := builder.BuildServices()
	err := builder.ServeGRPCServer(services)
	if err != nil {
		log.Fatalln(err)
	}
}
