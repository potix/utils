package main

import (
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

func (t *testPlugin) Call(request interface{}) (interface{}, error) {
	return "hello " + request.(string), nil
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

