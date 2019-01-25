package main

import (
	"filter-service/config"
	"filter-service/service"
	"flag"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		panic(err)
	}

	return
}
