package main

import (
	"github.com/potix/utils/plugger"
)

type test struct {
	caller     *plugger.Caller
	configFile string
}

func (h *httpWatcher) Initialize() (error) {
	return nil
}

func (h *httpWatcher) Finalize() {
}

func (h *httpWatcher) Call(request interface{}) (interface{}, error) {
	return "hello", nil
}

func newTest(caller *plugger.Caller, configFile string) (plugger.Plugin, error) {
	return &test{
		caller: caller,
		configfile: confFile,
	}, nil
}

func GetPluginInfo() (string, plugger.PluginNewFunc) {
	return "test", newTestFunc
}
