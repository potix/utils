package configurator

import (
	"fmt"
	"log"
	"github.com/BurntSushi/toml"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type fileLoader struct {
	configFile string
}

func (f *fileLoader) loadToml(config interface{}) (error) {
	_, err := toml.DecodeFile(f.configFile, config)
	if err != nil {
		return fmt.Errorf("can not decode config file for toml (%v): %w", f.configFile, err)
	}
	return nil
}

func (f *fileLoader) loadJSON(config interface{}) (error) {
	buf, err := ioutil.ReadFile(f.configFile)
	if err != nil {
		return fmt.Errorf("can not read config file (%v): %w", f.configFile, err)
	}
	err = json.Unmarshal(buf, config)
	if err != nil {
		return fmt.Errorf("can not decode config file for json (%v): %w", f.configFile, err)
	}
	return nil
}

func (f *fileLoader) loadYaml(config interface{}) (error) {
	buf, err := ioutil.ReadFile(f.configFile)
	if err != nil {
		return fmt.Errorf("can not read config file (%v): %w", f.configFile, err)
	}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return fmt.Errorf("can not decode config file for yaml (%v): %w", f.configFile, err)
	}
	return nil
}

func (f *fileLoader)load(config interface{}) (error) {
	err := f.loadToml(config)
	if err == nil {
		log.Printf("load config file as toml (%v)", f.configFile)
		return nil
	}
	err = f.loadJSON(config)
	if err == nil {
		log.Printf("load config file as json (%v)", f.configFile)
		return nil
	}
	err = f.loadYaml(config)
	if err == nil {
		log.Printf("load config file as yaml (%v)", f.configFile)
		return nil
	}
	return fmt.Errorf("can not decode config file (%v): %w", f.configFile, err)
}

func newFileLoader(configFile string) (*fileLoader) {
	return &fileLoader{
            configFile: configFile,
        }
}
