package goodfee

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

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

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().UnixNano()
	fee := npool.GoodFee{
		GoodID:         uuid.New().String(),
		AppID:          uuid.New().String(),
		FeeType:        fmt.Sprintf("GasFee-%v", nano),
		FeeDescription: "jkjdsajlkfdlsajfdlksajlkfdjsal;fjdsa",
		PayType:        "amount",
	}

	resp, err := Create(context.Background(), &npool.CreateGoodFeeRequest{
		Info: &fee,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.Equal(t, resp.Info.GoodID, fee.GoodID)
		assert.Equal(t, resp.Info.AppID, fee.AppID)
		assert.Equal(t, resp.Info.FeeType, fee.FeeType)
		assert.Equal(t, resp.Info.FeeDescription, fee.FeeDescription)
		assert.Equal(t, resp.Info.PayType, fee.PayType)
	}

	resp1, err := Get(context.Background(), &npool.GetGoodFeeRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp.Info.GoodID, fee.GoodID)
		assert.Equal(t, resp.Info.AppID, fee.AppID)
		assert.Equal(t, resp.Info.FeeType, fee.FeeType)
		assert.Equal(t, resp.Info.FeeDescription, fee.FeeDescription)
		assert.Equal(t, resp.Info.PayType, fee.PayType)
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	resp, err := GetAll(context.Background(), &npool.GetGoodFeesRequest{})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
