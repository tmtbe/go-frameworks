package main

import "flag"

var configFile = flag.String("f", "resources/configs/application.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()
	application, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}
	if err := application.Start(); err != nil {
		panic(err)
	}
	application.AwaitSignal()
}
