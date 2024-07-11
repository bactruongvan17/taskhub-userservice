package main

import (
	"log"
	"os"

	"github.com/bactruongvan17/taskhub-userservice/src/conf"
	"github.com/bactruongvan17/taskhub-userservice/src/route"
)

func main() {
	// load config
	err := conf.SetEnv()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	app := route.NewService()
	app.Run()
}
