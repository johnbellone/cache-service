package cache

import (
	"context"
	//timestamp "github.com/golang/protobuf/ptypes/timestamp"
	//wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/cache-service/proto/cache"
	//log "github.com/micro/go-micro/v2/logger"
)

type Handler struct{}

func (h *Handler) List(ctx context.Context, in *pb.CacheListRequest, out *pb.CacheListResponse) error {
	return nil
}

func (h *Handler) Put(ctx context.Context, in *pb.CacheMessage, out *pb.CacheMessage) error {
	out.Key = in.Key
	return nil
}

func (h *Handler) Set(ctx context.Context, in *pb.CacheMessage, out *pb.CacheMessage) error {
	out.Key = in.Key
	return nil
}

func (h *Handler) Get(ctx context.Context, in *pb.CacheKeyRequest, out *pb.CacheMessage) error {
	out.Key = in.Key
	return nil
}

func (h *Handler) Delete(ctx context.Context, in *pb.CacheKeyRequest, out *pb.CacheMessage) error {
	out.Key = in.Key
	return nil
}
