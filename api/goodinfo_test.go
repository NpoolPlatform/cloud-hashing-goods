// +build !codeanalysis

package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
)

func assertGoodInfo(t *testing.T, actual, expected *npool.GoodInfo) {
	assert.Equal(t, actual.DeviceInfoID, expected.DeviceInfoID)
	assert.Equal(t, actual.SeparateFee, expected.SeparateFee)
	assert.Equal(t, actual.UnitPower, expected.UnitPower)
	assert.Equal(t, actual.DurationDays, expected.DurationDays)
	assert.Equal(t, actual.CoinInfoID, expected.CoinInfoID)
	assert.Equal(t, actual.DeliveryAt, expected.DeliveryAt)
	assert.Equal(t, actual.Actuals, expected.Actuals)
	assert.Equal(t, actual.InheritFromGoodID, expected.InheritFromGoodID)
	assert.Equal(t, actual.VendorLocationID, expected.VendorLocationID)
	assert.Equal(t, actual.Price, expected.Price)
	assert.Equal(t, actual.BenefitType, expected.BenefitType)
	assert.Equal(t, actual.Classic, expected.Classic)
	assert.Equal(t, actual.SupportCoinTypeIDs, expected.SupportCoinTypeIDs)
}

func TestGoodCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	deviceInfoID := uuid.New().String()
	separateFee := true
	unitPower := int32(100)
	duration := int32(180)
	coinInfoID := uuid.New().String()
	actuals := true
	deliveryAt := uint32(time.Now().Unix())
	inheritFromGoodID := uuid.New().String()
	vendorLocationID := uuid.New().String()
	price := 100.8
	benefitType := "pool"
	classic := true
	supportCoinTypeIDs := []string{uuid.New().String(), uuid.New().String()}

	goodInfo := npool.GoodInfo{
		DeviceInfoID:       deviceInfoID,
		SeparateFee:        separateFee,
		UnitPower:          unitPower,
		DurationDays:       duration,
		CoinInfoID:         coinInfoID,
		DeliveryAt:         deliveryAt,
		Actuals:            actuals,
		InheritFromGoodID:  inheritFromGoodID,
		VendorLocationID:   vendorLocationID,
		Price:              price,
		PriceCurrency:      uuid.New().String(),
		BenefitType:        benefitType,
		Classic:            classic,
		SupportCoinTypeIDs: supportCoinTypeIDs,
		Title:              "Ant Miner S19 Pro",
		Unit:               "TH/s",
	}
	firstCreateInfo := npool.CreateGoodResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:50020/v1/create/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assertGoodInfo(t, firstCreateInfo.Info, &goodInfo)
		}
	}

	goodInfo.BenefitType = "platform"
	goodInfo.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:50020/v1/update/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateGoodResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assertGoodInfo(t, info.Info, &goodInfo)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodRequest{
			ID: goodInfo.ID,
		}).
		Post("http://localhost:50020/v1/get/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assertGoodInfo(t, info.Info, &goodInfo)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodRequest{
			ID: goodInfo.ID,
		}).
		Post("http://localhost:50020/v1/delete/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteGoodResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assertGoodInfo(t, info.Info, &goodInfo)
		}
	}

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:50020/v1/get/goods")
	assert.Nil(t, err)
}
