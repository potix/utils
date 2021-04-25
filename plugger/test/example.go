package main

import (
	"log"
        "github.com/potix/utils/plugger"
)

func main() {
	err := plugger.LoadPlugins("./plugin")
	if err != nil {
		log.Fatalf("can not load plugins: %v", err)
	}
}
