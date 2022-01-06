package tests

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/internal/app/module1/domain/services"
	"test/tests/mocks"
	"testing"
)

func TestUserDetailApi(t *testing.T) {
	flag.Parse()
	us := new(mocks.UserDetailService)
	us.On("GetUserDetail", mock.AnythingOfType("uint64")).Return(&services.UserDetail{}, nil)
	api, err := CreateUserDetailAPI(*resourcesPath, us)
	if err != nil {
		t.Fatalf("get userDetail api  error,%+v", err)
	}
	resp := callAPI(api.API, "GET", "/detail?id=1", nil)
	assert.Equal(t, 200, resp.StatusCode)
}
