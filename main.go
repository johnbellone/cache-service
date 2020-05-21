package main

import (
	"time"

	"github.com/johnbellone/cache-service/cache"
	pb "github.com/johnbellone/cache-service/protobuf"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	service := micro.NewService(
		micro.Name("cache-service"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	pb.RegisterCacheHandler(service.Server(), new(cache.Handler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
