package configurator

import (
	"fmt"
	"os"
)

// Configurator is configrator
type Configurator struct {
	loader *fileLoader
}

// Load is load
func (c *Configurator) Load(config interface{}) error {
	err := c.loader.load(config)
	return err
}

func validateConfigFile(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("can not open config file (%v): %w", configFile, err)
	}
	f.Close()
	return nil
}

// NewConfigurator is create new configurator
func NewConfigurator(configFile string) (*Configurator, error) {
	err := validateConfigFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("invalid config file (%v): %w", configFile, err)
	}
	newConfigurator := &Configurator{
		loader: newFileLoader(configFile),
	}
	return newConfigurator, nil
}
