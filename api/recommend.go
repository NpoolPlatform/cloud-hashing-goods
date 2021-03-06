// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	crud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-goods/pkg/middleware/recommend" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create recommend good: %v", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateRecommendForOtherApp(ctx context.Context, in *npool.CreateRecommendForOtherAppRequest) (*npool.CreateRecommendForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateRecommendRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create recommend good: %v", err)
		return &npool.CreateRecommendForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateRecommendForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateRecommend(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.UpdateRecommendResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update recommend good: %v", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetRecommendsByApp(ctx context.Context, in *npool.GetRecommendsByAppRequest) (*npool.GetRecommendsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods by app: %v", err)
		return &npool.GetRecommendsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetRecommendsByOtherApp(ctx context.Context, in *npool.GetRecommendsByOtherAppRequest) (*npool.GetRecommendsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetRecommendsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get recommend goods by app: %v", err)
		return &npool.GetRecommendsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetRecommendsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetRecommendsByRecommender(ctx context.Context, in *npool.GetRecommendsByRecommenderRequest) (*npool.GetRecommendsByRecommenderResponse, error) {
	resp, err := crud.GetByRecommender(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods by recommender: %v", err)
		return &npool.GetRecommendsByRecommenderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteRecommend(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	resp, err := crud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete recommend good: %v", err)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetRecommendGoodsByApp(ctx context.Context, in *npool.GetRecommendGoodsByAppRequest) (*npool.GetRecommendGoodsByAppResponse, error) {
	resp, err := mw.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods by app: %v", err)
		return &npool.GetRecommendGoodsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetRecommendGoodsByOtherApp(ctx context.Context, in *npool.GetRecommendGoodsByOtherAppRequest) (*npool.GetRecommendGoodsByOtherAppResponse, error) {
	resp, err := mw.GetByApp(ctx, &npool.GetRecommendGoodsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get recommend goods by app: %v", err)
		return &npool.GetRecommendGoodsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetRecommendGoodsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetRecommendGoodsByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	resp, err := mw.GetByRecommender(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods by app: %v", err)
		return &npool.GetRecommendGoodsByRecommenderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
