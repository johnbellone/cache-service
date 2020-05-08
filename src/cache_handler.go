package main

import (
	"context"
	//"github.com/micro/go-micro/v2"
	//"github.com/micro/go-micro/v2/auth"
	//"github.com/micro/go-micro/v2/errors"
	pb "github.com/johnbellone/cache-service/api"
)

type CacheHandler struct{}

func (h *CacheHandler) Get(ctx context.Context, req *pb.CacheRequest, rsp *pb.CacheResponse) error {
	return nil
}

func (h *CacheHandler) Put(ctx context.Context, req *pb.CacheRequest, rsp *pb.CacheResponse) error {
	return nil
}
