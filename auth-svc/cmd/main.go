package main

import (
	"light-ms/auth-svc/internal"
	"log"
)

func main() {
	router := internal.NewRouter(internal.LoadConf())
	router.Register()
	if e := router.Run(); e != nil {
		log.Fatal(e)
	}
}
