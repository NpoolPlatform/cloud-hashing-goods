package targetarea

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	db "github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/service-name" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/app"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)

	err := app.Init(servicename.ServiceName, "", "", "", configPath, nil, nil)
	if err != nil {
		return
	}
	err = db.Init()
	if err != nil {
		return
	}
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	continent := "AsiaPackageApiTest"
	country := "ChinaPackageApiTest"
	country1 := "ChinaPackageApiTest"

	resp, err := Create(context.Background(), &npool.CreateTargetAreaRequest{
		Info: &npool.TargetAreaInfo{
			Continent: continent,
			Country:   country,
		},
	})
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		if assert.NotNil(t, resp.Info) {
			assert.Equal(t, resp.Info.Continent, continent)
			assert.Equal(t, resp.Info.Country, country)
		}
	}

	_, err = Create(context.Background(), &npool.CreateTargetAreaRequest{
		Info: &npool.TargetAreaInfo{
			Continent: continent,
			Country:   country,
		},
	})
	assert.NotNil(t, err)

	resp.Info.Country = country1

	resp1, err := Update(context.Background(), &npool.UpdateTargetAreaRequest{
		Info: resp.Info,
	})
	assert.Nil(t, err)
	if assert.NotNil(t, resp1) {
		if assert.NotNil(t, resp1.Info) {
			assert.Equal(t, resp1.Info.ID, resp.Info.ID)
			assert.Equal(t, resp1.Info.Continent, continent)
			assert.Equal(t, resp1.Info.Country, country1)
		}
	}

	resp2, err := Delete(context.Background(), &npool.DeleteTargetAreaRequest{
		ID: resp1.Info.ID,
	})
	if assert.Nil(t, err) {
		if assert.NotNil(t, resp2.Info) {
			assert.Equal(t, resp2.Info.ID, resp1.Info.ID)
			assert.Equal(t, resp2.Info.Continent, continent)
			assert.Equal(t, resp2.Info.Country, country1)
		}
	}

	resp, err = Create(context.Background(), &npool.CreateTargetAreaRequest{
		Info: &npool.TargetAreaInfo{
			Continent: continent,
			Country:   country,
		},
	})
	assert.Nil(t, err)

	resp3, err := DeleteByContinentCountry(context.Background(), &npool.DeleteTargetAreaByContinentCountryRequest{
		Continent: continent,
		Country:   country,
	})
	if assert.Nil(t, err) {
		if assert.NotNil(t, resp3.Info) {
			assert.Equal(t, resp3.Info.ID, resp.Info.ID)
			assert.Equal(t, resp3.Info.Continent, continent)
			assert.Equal(t, resp3.Info.Country, country)
		}
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	resp, err := GetAll(context.Background(), &npool.GetTargetAreasRequest{})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
