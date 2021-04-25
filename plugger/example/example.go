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
        pluginCtx, err := plugger.GetPluginContext("testplugin", "caller", "config_file")
	if err != nil {
		log.Fatalf("can not get test plugin context: %v", err)
	}
	err = pluginCtx.Initialize()
	if err != nil {
		log.Fatalf("failed to plugin initialize: %v", err)
	}
	defer pluginCtx.Finalize()
	response, err := pluginCtx.Call("world")
	if err != nil {
		log.Fatalf("faild to plugin call: %v", err)
	}
	log.Println(response.(string))
}

