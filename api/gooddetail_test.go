package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
)

func assertGoodDetail(t *testing.T, actual *npool.GoodDetail, expectGoodInfo *npool.GoodInfo, expectDevice *npool.DeviceInfo, expectVendorLocation *npool.VendorLocationInfo) {
	assert.Equal(t, actual.DeviceInfo.ID, expectDevice.ID)
	assert.Equal(t, actual.DeviceInfo.Type, expectDevice.Type)
	assert.Equal(t, actual.DeviceInfo.Manufacturer, expectDevice.Manufacturer)
	assert.Equal(t, actual.DeviceInfo.PowerComsuption, expectDevice.PowerComsuption)
	assert.Equal(t, actual.DeviceInfo.ShipmentAt, expectDevice.ShipmentAt)
	assert.Equal(t, actual.Good.SeparateFee, expectGoodInfo.SeparateFee)
	assert.Equal(t, actual.Good.UnitPower, expectGoodInfo.UnitPower)
	assert.Equal(t, actual.Good.DurationDays, expectGoodInfo.DurationDays)
	assert.Equal(t, actual.Good.CoinInfoID, expectGoodInfo.CoinInfoID)
	assert.Equal(t, actual.Good.Actuals, expectGoodInfo.Actuals)
	assert.Equal(t, actual.Good.DeliveryAt, expectGoodInfo.DeliveryAt)
	assert.Equal(t, actual.VendorLocation.ID, expectVendorLocation.ID)
	assert.Equal(t, actual.VendorLocation.Country, expectVendorLocation.Country)
	assert.Equal(t, actual.VendorLocation.Province, expectVendorLocation.Province)
	assert.Equal(t, actual.VendorLocation.City, expectVendorLocation.City)
	assert.Equal(t, actual.VendorLocation.Address, expectVendorLocation.Address)
	assert.Equal(t, actual.Good.Price, expectGoodInfo.Price)
	assert.Equal(t, actual.Good.BenefitType, expectGoodInfo.BenefitType)
	assert.Equal(t, actual.Good.Classic, expectGoodInfo.Classic)
	assert.Equal(t, actual.Good.SupportCoinTypeIDs, expectGoodInfo.SupportCoinTypeIDs)
}

func TestGet(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	nano := time.Now().Unix()

	deviceInfo := npool.DeviceInfo{
		Type:            fmt.Sprintf("S19-%v", nano),
		Manufacturer:    fmt.Sprintf("Ant-%v", nano),
		PowerComsuption: 120,
		ShipmentAt:      int32(time.Now().Unix()),
	}

	deviceResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateDeviceInfoRequest{
			Info: &deviceInfo,
		}).
		Post("http://localhost:50020/v1/create/device")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, deviceResp.StatusCode())
		info := npool.CreateDeviceInfoResponse{}
		err := json.Unmarshal(deviceResp.Body(), &info)
		if assert.Nil(t, err) {
			deviceInfo.ID = info.Info.ID
		}
	}

	vendorLocation := npool.VendorLocationInfo{
		Country:  fmt.Sprintf("China-%v", nano),
		Province: fmt.Sprintf("Shanghai-%v", nano),
		City:     fmt.Sprintf("Shanghai-%v", nano),
		Address:  fmt.Sprintf("Shanghai-%v", nano),
	}

	vendorLocationResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateVendorLocationRequest{
			Info: &vendorLocation,
		}).
		Post("http://localhost:50020/v1/create/vendor-location")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, vendorLocationResp.StatusCode())
		info := npool.CreateVendorLocationResponse{}
		err := json.Unmarshal(vendorLocationResp.Body(), &info)
		if assert.Nil(t, err) {
			vendorLocation.ID = info.Info.ID
		}
	}

	currency := npool.PriceCurrency{
		Name:   fmt.Sprintf("USDT-%v", nano),
		Unit:   "USDT",
		Symbol: "$",
	}
	currencyResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreatePriceCurrencyRequest{
			Info: &currency,
		}).
		Post("http://localhost:50020/v1/create/price-currency")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, currencyResp.StatusCode())
		info := npool.CreatePriceCurrencyResponse{}
		err := json.Unmarshal(currencyResp.Body(), &info)
		if assert.Nil(t, err) {
			currency.ID = info.Info.ID
		}
	}

	goodInfo := npool.GoodInfo{
		DeviceInfoID:       deviceInfo.ID,
		SeparateFee:        true,
		UnitPower:          10,
		DurationDays:       180,
		CoinInfoID:         uuid.New().String(),
		Actuals:            true,
		DeliveryAt:         uint32(time.Now().Unix()),
		VendorLocationID:   vendorLocation.ID,
		InheritFromGoodID:  uuid.UUID{}.String(),
		Price:              13.0,
		PriceCurrency:      currency.ID,
		BenefitType:        "platform",
		Classic:            true,
		SupportCoinTypeIDs: []string{uuid.New().String(), uuid.New().String()},
		Title:              "Ant Miner S19 Pro",
		Unit:               "TH/s",
	}

	goodInfoResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:50020/v1/create/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, goodInfoResp.StatusCode())
		info := npool.CreateGoodResponse{}
		err := json.Unmarshal(goodInfoResp.Body(), &info)
		if assert.Nil(t, err) {
			goodInfo.ID = info.Info.ID
		}
	}

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodDetailRequest{
			ID: goodInfo.ID,
		}).
		Post("http://localhost:50020/v1/get/good/detail")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodDetailResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.Good.ID, goodInfo.ID)
			assert.Nil(t, info.Info.InheritFromGood)
			assertGoodDetail(t, info.Info, &goodInfo, &deviceInfo, &vendorLocation)
		}
	}

	inheritGoodID := goodInfo.ID
	goodInfo.InheritFromGoodID = goodInfo.ID
	goodInfoResp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:50020/v1/create/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, goodInfoResp1.StatusCode())
		info := npool.CreateGoodResponse{}
		err := json.Unmarshal(goodInfoResp1.Body(), &info)
		if assert.Nil(t, err) {
			goodInfo.ID = info.Info.ID
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodDetailRequest{
			ID: goodInfo.ID,
		}).
		Post("http://localhost:50020/v1/get/good/detail")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodDetailResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.Good.ID, goodInfo.ID)
			assert.Equal(t, info.Info.InheritFromGood.ID, inheritGoodID)
			assertGoodDetail(t, info.Info, &goodInfo, &deviceInfo, &vendorLocation)
		}
	}
}
