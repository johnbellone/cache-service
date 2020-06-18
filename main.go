//go:generate protoc -Iproto --micro_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto proto/cache/cache_service.proto
package main

import (
	"time"

	"github.com/johnbellone/cache-service/internal/cache"
	pb "github.com/johnbellone/cache-service/proto/cache"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.cache"),
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
