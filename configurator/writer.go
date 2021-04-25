package configurator

import (
	"github.com/BurntSushi/toml"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"io/ioutil"
	"bytes"
	"fmt"
	"os"
)

type fileWriter struct {
	configFile string
}

func (w *fileWriter) write(formatType FormatType, data interface{}) (error) {
	tmpConfigFile := w.configFile + ".tmp"
	if formatType == FormatTypeAuto {
		ext := filepath.Ext(w.configFile)
		switch ext {
		case ".tml":
			fallthrough
		case ".toml":
			formatType = FormatTypeToml
		case ".yml":
			fallthrough
		case ".yaml":
			formatType = FormatTypeYaml
		case ".jsn":
			fallthrough
		case ".json":
			formatType = FormatTypeJson
		default:
			return fmt.Errorf("unexpected file extension (ext = %v)", ext)
		}
	}
        switch formatType {
        case FormatTypeToml:
		var buffer bytes.Buffer
		encoder := toml.NewEncoder(&buffer)
		err := encoder.Encode(data)
		if err != nil {
			return fmt.Errorf("can not encode with toml (config = %v): %w", tmpConfigFile, err)
		}
		err = ioutil.WriteFile(tmpConfigFile, buffer.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("can not write file with toml (config = %v): %w", tmpConfigFile, err)
		}
        case FormatTypeYaml:
		y, err := yaml.Marshal(data)
                if err != nil {
			return fmt.Errorf("can not encode with yaml (cofnig = %v): %w", tmpConfigFile, err)
                }
		err = ioutil.WriteFile(tmpConfigFile, y, 0644)
		if err != nil {
			return fmt.Errorf("can not write file with yaml (config = %v): %w", tmpConfigFile, err)
		}
        case FormatTypeJson:
		j, err := json.Marshal(data)
                if err != nil {
			return fmt.Errorf("can not encode with json (config = %v): %w", tmpConfigFile, err)
                }
		if err = ioutil.WriteFile(tmpConfigFile, j, 0644);  err != nil {
			return fmt.Errorf("can not write file with json (config = %v): %w", tmpConfigFile, err)
		}
	default:
		return fmt.Errorf("unsupported format type (formatType = %v)", formatType)
        }
	if err := os.Rename(tmpConfigFile, w.configFile); err != nil {
		return fmt.Errorf("can not rename file (src = %v, dst = %v): %w", tmpConfigFile, w.configFile, err)
	}
	return nil
}

func newFileWriter(configFile string) (*fileWriter) {
	return &fileWriter{
		configFile: configFile,
	}
}
