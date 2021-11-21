// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/middleware/good-detail" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetGoodDetail(ctx context.Context, in *npool.GetGoodDetailRequest) (*npool.GetGoodDetailResponse, error) {
	resp, err := gooddetail.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good detail error: %w", err)
		return &npool.GetGoodDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodsDetail(ctx context.Context, in *npool.GetGoodsDetailRequest) (*npool.GetGoodsDetailResponse, error) {
	resp, err := gooddetail.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good detail all error: %w", err)
		return &npool.GetGoodsDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
