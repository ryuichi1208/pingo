package service

import (
    "context"
    pb "github.com/TakeruTakeru/gserver/pb"
)

type IncrementService struct {
    cacheNum int32
}

func (s *IncrementService) Increment(ctx context.Context, req *pb.IncrementRequest) (*pb.IncrementResponse, error) {
    n := req.GetNumber() + 1
    return &pb.IncrementResponse{Number: n}, nil
}

func (s *IncrementService) GetAndIncrement(ctx context.Context, req *pb.IncrementRequest) (*pb.IncrementResponse, error) {
    if s.cacheNum != 0 {
        s.cacheNum = s.cacheNum + 1
    } else {
        s.cacheNum = req.GetNumber()
    }
    return &pb.IncrementResponse{Number: s.cacheNum}, nil
}

func NewIncrementService() *IncrementService {
    return &IncrementService{}
}
