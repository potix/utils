package main

import (
	"github.com/potix/utils/plugger"
)

type test struct {
	caller     *Caller
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

func newTest(caller *Caller, configFile string) (Plugin, error) {
	return &test{
		caller: caller,
		configfile: confFile,
	}, nil
}

func GetPluginInfo() (string, PluginNewFunc) {
	return "test", newTestFunc
}
