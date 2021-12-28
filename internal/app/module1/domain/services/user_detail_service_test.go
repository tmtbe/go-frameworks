package services

import (
	"flag"
	"testing"
)

var configFile = flag.String("f", "resources/configs/application.yml", "set config file which viper will loading.")

func TestDetailsRepository_Get(t *testing.T) {
	flag.Parse()
	service, err := CreateUserDetailService(*configFile)
	if err != nil {
		return
	}
	service.GetUserDetail(1)
}
