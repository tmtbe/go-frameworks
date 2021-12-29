package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetailsRepository_GetUserDetail(t *testing.T) {
	background := setUp()
	background.MustSetUpDb("get_user_detail")
	detail, err := background.UserDetailService.GetUserDetail(1)
	if err != nil {
		t.Fatalf("userDetail service get userDetail error,%+v", err)
	}
	assert.Equal(t, uint64(1), detail.ID)
	assert.Equal(t, float32(1), detail.Price)
}
