package main

import (
	"flag"
	_ "github.com/lib/pq"
)

var resourcesPath = flag.String("f", ".", "set resources path viper will loading.")

func main() {
	flag.Parse()
	application, err := CreateApp(*resourcesPath)
	if err != nil {
		panic(err)
	}
	if err := application.Start(); err != nil {
		panic(err)
	}
	application.AwaitSignal()
}
