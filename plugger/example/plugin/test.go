package main

import (
	"log"
	"github.com/potix/utils/plugger"
)

type testPlugin struct {
	caller     *plugger.Caller
	configFile string
}

func (t *testPlugin) Initialize() (error) {
	return nil
}

func (t *testPlugin) Finalize() {
}

func (t *testPlugin) Call(methodName string, request interface{}) (interface{}, error) {
	if methodName == "whoAreYou" {
		log.Printf("%v: I am bob.", request.(int))
		response, err := t.caller.Callback("whoAreYou", request.(int) + 1)
		if err != nil {
			log.Fatalf("callback error: %v", err)
		}
		log.Println(response.(string))
	}
	return "bob ok!", nil
}

func newTest(caller *plugger.Caller, configFile string) (plugger.Plugin, error) {
	return &testPlugin{
		caller: caller,
		configFile: configFile,
	}, nil
}

func GetPluginInfo() (string, plugger.PluginNewFunc) {
	return "testplugin", newTest
}

