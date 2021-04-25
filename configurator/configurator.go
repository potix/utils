package configurator

import (
	"fmt"
	"os"
)

type FormatType int

const (
     FormatTypeNone FormatType = 0 + iota
     FormatTypeAuto
     FormatTypeToml
     FormatTypeYaml
     FormatTypeJson
)

// Configurator is configrator
type Configurator struct {
	formatType FormatType
	reader     *fileReader
	writer     *fileWriter
}

// Load is read
func (c *Configurator) Load(config interface{}) error {
	formatType, err := c.reader.read(config)
	if err != nil {
		return err
	}
	c.formatType = formatType
	return nil
}

func (c *Configurator) Save(formatType FormatType, config interface{}) (err error) {
	if formatType >= FormatTypeToml && formatType <= FormatTypeJson {
		return c.writer.write(formatType, config)
	}
	return c.writer.write(c.formatType, config)
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
		formatType: FormatTypeAuto,
		reader:     newFileReader(configFile),
		writer:     newFileWriter(configFile),
	}
	return newConfigurator, nil
}
