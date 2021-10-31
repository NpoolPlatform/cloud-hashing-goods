package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/target-area" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTargetArea(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	resp, err := targetarea.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create target area error: %w", err)
		return &npool.CreateTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateTargetArea(ctx context.Context, in *npool.UpdateTargetAreaRequest) (*npool.UpdateTargetAreaResponse, error) {
	resp, err := targetarea.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.UpdateTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetTargetAreas(ctx context.Context, in *npool.GetTargetAreasRequest) (*npool.GetTargetAreasResponse, error) {
	resp, err := targetarea.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get target areas error: %w", err)
		return &npool.GetTargetAreasResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}