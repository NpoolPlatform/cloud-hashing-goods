package appgood

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppGood(info *npool.AppGoodInfo) error {
	_, err := uuid.Parse(info.GetAppID())
	if err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	_, err = uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}

	return nil
}

func dbRowToAppGood(info *ent.AppGood) *npool.AppGoodInfo {
	return &npool.AppGoodInfo{
		ID:               info.ID.String(),
		AppID:            info.AppID.String(),
		GoodID:           info.GoodID.String(),
		Authorized:       info.Authorized,
		Online:           info.Online,
		InitAreaStrategy: string(info.InitAreaStrategy),
		Price:            price.DBPriceToVisualPrice(info.Price),
		GasPrice:         price.DBPriceToVisualPrice(info.GasPrice),
	}
}

func Authorize(ctx context.Context, in *npool.AuthorizeAppGoodRequest) (*npool.AuthorizeAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err == nil {
		info, err := db.Client().
			AppGood.
			UpdateOneID(id).
			SetAuthorized(true).
			SetDeleteAt(0).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail authorize app good: %v", err)
		}
		return &npool.AuthorizeAppGoodResponse{
			Info: dbRowToAppGood(info),
		}, nil
	}

	info, err := db.Client().
		AppGood.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetAuthorized(true).
		SetOnline(false).
		SetInitAreaStrategy(appgood.InitAreaStrategy(in.GetInfo().GetInitAreaStrategy())).
		SetPrice(0).
		SetGasPrice(0).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app good: %v", err)
	}

	return &npool.AuthorizeAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Check(ctx context.Context, in *npool.CheckAppGoodRequest) (*npool.CheckAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	infos, err := db.Client().
		AppGood.
		Query().
		Where(
			appgood.And(
				appgood.AppID(uuid.MustParse(in.GetInfo().GetAppID())),
				appgood.GoodID(uuid.MustParse(in.GetInfo().GetGoodID())),
				appgood.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app good: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty app good")
	}

	return &npool.CheckAppGoodResponse{
		Info: dbRowToAppGood(infos[0]),
	}, nil
}

func SetAppGoodPrice(ctx context.Context, in *npool.SetAppGoodPriceRequest) (*npool.SetAppGoodPriceResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	info1, err := Check(ctx, &npool.CheckAppGoodRequest{
		Info: in.GetInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail check app good: %v", err)
	}

	if !info1.Info.Authorized {
		return nil, xerrors.Errorf("good not authorize to app")
	}
	if info1.Info.Online {
		return nil, xerrors.Errorf("cannot set price to online good")
	}

	if price.VisualPriceToDBPrice(in.GetInfo().GetPrice()) == 0 ||
		price.VisualPriceToDBPrice(in.GetInfo().GetGasPrice()) == 0 {
		return nil, xerrors.Errorf("price should be greater than 0")
	}

	info, err := db.Client().
		AppGood.
		UpdateOneID(id).
		SetPrice(price.VisualPriceToDBPrice(in.GetInfo().GetPrice())).
		SetGasPrice(price.VisualPriceToDBPrice(in.GetInfo().GetGasPrice())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail set price app good: %v", err)
	}

	return &npool.SetAppGoodPriceResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Onsale(ctx context.Context, in *npool.OnsaleAppGoodRequest) (*npool.OnsaleAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	info1, err := Check(ctx, &npool.CheckAppGoodRequest{
		Info: in.GetInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail onsale app good: %v", err)
	}

	if !info1.Info.Authorized {
		return nil, xerrors.Errorf("good not authorized by app")
	}

	info, err := db.Client().
		AppGood.
		UpdateOneID(id).
		SetOnline(true).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail onsale app good: %v", err)
	}

	return &npool.OnsaleAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Offsale(ctx context.Context, in *npool.OffsaleAppGoodRequest) (*npool.OffsaleAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	info, err := db.Client().
		AppGood.
		UpdateOneID(id).
		SetOnline(false).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail offsale app good: %v", err)
	}

	return &npool.OffsaleAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Unauthorize(ctx context.Context, in *npool.UnauthorizeAppGoodRequest) (*npool.UnauthorizeAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	info, err := db.Client().
		AppGood.
		UpdateOneID(id).
		SetAuthorized(false).
		SetOnline(false).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail unauthorize app good: %v", err)
	}

	return &npool.UnauthorizeAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}