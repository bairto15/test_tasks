package main

import (
	"net/http"
	"runtime/debug"
	"test_puzzle/config"
	"test_puzzle/package/handler/http_net"
	"test_puzzle/package/logging"
	"test_puzzle/package/repository"
	"test_puzzle/package/service"

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

	postgres := repository.NewPostgresDB(conf)

	repository := repository.New(postgres)
	service := service.New(repository)
	handler := http_net.New(service)
	middleware := http_net.NewMiddleware(handler)

	http.ListenAndServe(conf.Host+conf.Port, middleware)	
}
