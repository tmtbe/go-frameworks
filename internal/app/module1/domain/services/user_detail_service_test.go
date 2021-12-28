package services

import (
	"flag"
	"testing"
)

var resourcesPath = flag.String("f", ".", "set resources file which viper will loading.")

func setUp() *TestContext {
	flag.Parse()
	testContext, err := CreateTestContext(*resourcesPath)
	if err != nil {
		panic(err)
	}
	return testContext
}

func TestDetailsRepository_GetUserDetail(t *testing.T) {
	testContext := setUp()
	testContext.userDetailService.GetUserDetail(1)
}
