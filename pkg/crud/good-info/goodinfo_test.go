// +build !codeanalysis

package goodinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestGoodInfoCRUD(t *testing.T) { //nolint
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

	resp, err := Create(context.Background(), &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp.Info.SeparateFee, separateFee)
		assert.Equal(t, resp.Info.UnitPower, unitPower)
		assert.Equal(t, resp.Info.DurationDays, duration)
		assert.Equal(t, resp.Info.CoinInfoID, coinInfoID)
		assert.Equal(t, resp.Info.Actuals, actuals)
		assert.Equal(t, resp.Info.InheritFromGoodID, inheritFromGoodID)
		assert.Equal(t, resp.Info.VendorLocationID, vendorLocationID)
		assert.Equal(t, resp.Info.Price, price)
		assert.Equal(t, resp.Info.BenefitType, benefitType)
		assert.Equal(t, resp.Info.Classic, classic)
		assert.Equal(t, resp.Info.SupportCoinTypeIDs, supportCoinTypeIDs)
		assert.Equal(t, resp.Info.PriceCurrency, goodInfo.PriceCurrency)
		assert.Equal(t, resp.Info.Title, goodInfo.Title)
		assert.Equal(t, resp.Info.Unit, goodInfo.Unit)
	}

	goodInfo.BenefitType = "platform"
	goodInfo.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodRequest{
		Info: &goodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp1.Info.SeparateFee, separateFee)
		assert.Equal(t, resp1.Info.UnitPower, unitPower)
		assert.Equal(t, resp1.Info.DurationDays, duration)
		assert.Equal(t, resp1.Info.CoinInfoID, coinInfoID)
		assert.Equal(t, resp1.Info.Actuals, actuals)
		assert.Equal(t, resp1.Info.InheritFromGoodID, inheritFromGoodID)
		assert.Equal(t, resp1.Info.VendorLocationID, vendorLocationID)
		assert.Equal(t, resp1.Info.Price, price)
		assert.Equal(t, resp1.Info.BenefitType, goodInfo.BenefitType)
		assert.Equal(t, resp1.Info.Classic, classic)
		assert.Equal(t, resp1.Info.SupportCoinTypeIDs, supportCoinTypeIDs)
	}

	resp2, err := Get(context.Background(), &npool.GetGoodRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assert.Equal(t, resp2.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp2.Info.SeparateFee, separateFee)
		assert.Equal(t, resp2.Info.UnitPower, unitPower)
		assert.Equal(t, resp2.Info.DurationDays, duration)
		assert.Equal(t, resp2.Info.CoinInfoID, coinInfoID)
		assert.Equal(t, resp2.Info.Actuals, actuals)
		assert.Equal(t, resp2.Info.InheritFromGoodID, inheritFromGoodID)
		assert.Equal(t, resp2.Info.VendorLocationID, vendorLocationID)
		assert.Equal(t, resp2.Info.Price, price)
		assert.Equal(t, resp2.Info.BenefitType, goodInfo.BenefitType)
		assert.Equal(t, resp2.Info.Classic, classic)
		assert.Equal(t, resp2.Info.SupportCoinTypeIDs, supportCoinTypeIDs)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteGoodRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assert.Equal(t, resp3.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp3.Info.SeparateFee, separateFee)
		assert.Equal(t, resp3.Info.UnitPower, unitPower)
		assert.Equal(t, resp3.Info.DurationDays, duration)
		assert.Equal(t, resp3.Info.CoinInfoID, coinInfoID)
		assert.Equal(t, resp3.Info.Actuals, actuals)
		assert.Equal(t, resp3.Info.InheritFromGoodID, inheritFromGoodID)
		assert.Equal(t, resp3.Info.VendorLocationID, vendorLocationID)
		assert.Equal(t, resp3.Info.Price, price)
		assert.Equal(t, resp3.Info.BenefitType, goodInfo.BenefitType)
		assert.Equal(t, resp3.Info.Classic, classic)
		assert.Equal(t, resp3.Info.SupportCoinTypeIDs, supportCoinTypeIDs)
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	_, err := GetAll(context.Background(), &npool.GetGoodsRequest{})
	assert.Nil(t, err)
}
