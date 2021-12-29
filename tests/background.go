package tests

import (
	"flag"
	_ "github.com/lib/pq"
	"test/tests/pkg/testcontainer"
)

var resourcesPath = flag.String("f", "../.", "set resources file which viper will loading.")

func setUp() *testcontainer.Background {
	flag.Parse()
	testContext, err := CreateBackground(*resourcesPath)
	if err != nil {
		panic(err)
	}
	return testContext
}
