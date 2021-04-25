package main

import (
        "github.com/potix/utils/plugger"
)

func main() {
	err := LoadPlugins("./plugin")
	if err != nil {
		t.Errorf("can not load plugins: %v", err)
	}
}
