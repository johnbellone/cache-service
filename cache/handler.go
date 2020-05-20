package cache

import (
	"context"

	pb "github.com/johnbellone/cache-service/protobuf"
)

type Handler struct{}

func (h *Handler) List(ctx context.Context, req *pb.CacheListRequest, rsp *pb.CacheListResponse) error {
	return nil
}

func (h *Handler) Put(ctx context.Context, req *pb.CacheMessage, rsp *pb.CacheMessage) error {
	return nil
}

func (h *Handler) Set(ctx context.Context, req *pb.CacheMessage, rsp *pb.CacheMessage) error {
	return nil
}

func (h *Handler) Get(ctx context.Context, req *pb.CacheMessage, rsp *pb.CacheMessage) error {
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.CacheMessage, rsp *pb.CacheMessage) error {
	return nil
}
