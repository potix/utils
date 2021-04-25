package main

import (
	"log"
        "github.com/potix/utils/plugger"
)

func whoAreyou(request interface{}) (interface{}, error) {
	log.Printf("%v: I am alice.", request.(int))
	return "alice ok!", nil
}

func main() {
	err := plugger.LoadPlugins("./plugin")
	if err != nil {
		log.Fatalf("can not load plugins: %v", err)
	}
        pluginCtx, err := plugger.GetPluginContext("testplugin", "caller", "config_file")
	if err != nil {
		log.Fatalf("can not get test plugin context: %v", err)
	}
	pluginCtx.SetCallback("whoAreYou", whoAreyou)
	err = pluginCtx.Initialize()
	if err != nil {
		log.Fatalf("failed to plugin initialize: %v", err)
	}
	defer pluginCtx.Finalize()
	response, err := pluginCtx.Call("whoAreYou", 1)
	if err != nil {
		log.Fatalf("faild to plugin call: %v", err)
	}
	log.Println(response.(string))
}

