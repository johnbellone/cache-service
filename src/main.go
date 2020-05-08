package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	pb "github.com/johnbellone/cache-service/api"
)

// Version The main, base version of the project.
const Version = "0.1.0"

// GitCommit The git commit of the build; this is set by the compiler.
var GitCommit string

// GitCommit The git commit of the build; this is set by the compiler.
var BuildDate string

func main() {
	service := micro.NewService(
		micro.Name("cache-service"),
		micro.Version(Version),
	)

	service.Init()
	pb.RegisterCacheHandler(service.Server(), new(CacheHandler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
