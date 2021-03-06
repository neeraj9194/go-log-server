package config

import (
	"io/ioutil"
	"fmt"
	"github.com/jinzhu/configor"
)

// Unmarshaler is the interface for objects that can unmarshal data into a target object
type Unmarshaler func(content []byte, target interface{}) error

// Loader loads the configuration from the configuration file into the configuration object
type Loader struct{ path string }

// NewLoader creates a new configuration loader object using the provided configuration path
func NewLoader(path string) *Loader {
	return &Loader{path}
}

// Load loads the configuration from the configuration file into the configuration object
func (cl *Loader) Load(unmarshaler Unmarshaler, target interface{}) error {
	content, err := ioutil.ReadFile(cl.path)
	if err != nil {
		return err
	}

	return unmarshaler(content, target)
}

type ServerConfig struct {
	Port	string
	Backend	string
	Root	string
}

var RootDir string

func LoadConfig(fileName string) ServerConfig {
	config := ServerConfig{}
	configor.Load(&config, fileName)
	fmt.Printf("Server config: %#v\n", config)
	RootDir = config.Root
	return config
}
