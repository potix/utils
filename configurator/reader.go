package configurator

import (
	"fmt"
	"log"
	"github.com/BurntSushi/toml"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type fileReader struct {
	configFile string
}

func (f *fileReader) readToml(config interface{}) (error) {
	_, err := toml.DecodeFile(f.configFile, config)
	if err != nil {
		return fmt.Errorf("can not decode config file for toml (%v): %w", f.configFile, err)
	}
	return nil
}

func (f *fileReader) readJSON(config interface{}) (error) {
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

func (f *fileReader) readYaml(config interface{}) (error) {
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

func (f *fileReader)read(config interface{}) (FormatType, error) {
	err := f.readToml(config)
	if err == nil {
		log.Printf("read config file as toml (%v)", f.configFile)
		return FormatTypeToml, nil
	}
	err = f.readJSON(config)
	if err == nil {
		log.Printf("read config file as json (%v)", f.configFile)
		return FormatTypeJson, nil
	}
	err = f.readYaml(config)
	if err == nil {
		log.Printf("read config file as yaml (%v)", f.configFile)
		return FormatTypeYaml, nil
	}
	return FormatTypeNone, fmt.Errorf("can not decode config file (%v): %w", f.configFile, err)
}

func newFileReader(configFile string) (*fileReader) {
	return &fileReader{
            configFile: configFile,
        }
}
