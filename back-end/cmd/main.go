package main

import (
	"net/http"
	"runtime/debug"
	"sync"
	"test_puzzle/config"
	"test_puzzle/package/handler/gin"
	"test_puzzle/package/handler/http_net"
	"test_puzzle/package/logging"
	"test_puzzle/package/repository"
	"test_puzzle/package/service"
	"test_puzzle/server"

	_ "github.com/lib/pq"
)

func main() {
	logger := logging.GetLogger()

	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("panic: %v,\n%s", r, debug.Stack())
		}
	}()

	conf, err := config.Get()
	if err != nil {
		logger.Fatal(err)
	}

	var mutex sync.Mutex 
	
	postgres := repository.NewPostgresDB(conf)

	repository := repository.New(postgres)
	service := service.New(repository, &mutex)
	handler := http_net.New(service)
	middleware := http_net.NewMiddleware(handler)

	handlers := gin.New(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(conf.PortServerGin, handlers.InitRoutes()); err != nil {
			logger.Fatalf("Error running http Server: %s", err.Error())
		}
	}()

	http.ListenAndServe(conf.Host+conf.PortServer, middleware)	
}
