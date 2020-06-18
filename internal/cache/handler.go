package cache

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/cache-service/proto/cache"
	log "github.com/micro/go-micro/v2/logger"
)

type Handler struct{}

func (h *Handler) List(ctx context.Context, req *pb.CacheListRequest, rsp *pb.CacheListResponse) error {
	log.Info("Received Cache.List request")
	return nil
}

func (h *Handler) Put(ctx context.Context, req *pb.CacheMessage, rsp *timestamp.Timestamp) error {
	log.Info("Received Cache.Put request")
	return nil
}

func (h *Handler) Set(ctx context.Context, req *pb.CacheMessage, rsp *timestamp.Timestamp) error {
	log.Info("Received Cache.Set request")
	return nil
}

func (h *Handler) Get(ctx context.Context, req *wrappers.BytesValue, rsp *pb.CacheMessage) error {
	log.Info("Received Cache.Get request")
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *wrappers.BytesValue, rsp *pb.CacheMessage) error {
	log.Info("Received Cache.Delete request")
	return nil
}
