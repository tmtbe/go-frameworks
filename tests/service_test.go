package tests

import (
	"testing"
)

func TestDetailsRepository_GetUserDetail(t *testing.T) {
	background := setUp()
	background.UserDetailService.GetUserDetail(1)
}
func TestDetailsRepository_GetUserDetail2(t *testing.T) {
	background := setUp()
	background.UserDetailService.GetUserDetail(1)
}
