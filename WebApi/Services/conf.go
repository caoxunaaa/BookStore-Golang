package Services

import (
	"gopkg.in/yaml.v2"
	"os"
)

var C *Config

type Config struct {
	Host    HostConfig    `yaml:"Host"`
	Jwt     JwtConfig     `yaml:"Jwt"`
	UserRpc UserRpcConfig `yaml:"UserRpc"`
}

type HostConfig struct {
	ListenOn string `yaml:"ListenOn"`
}

type UserRpcConfig struct {
	Host string `yaml:"Host"`
}

type JwtConfig struct {
	Secret string `yaml:"Secret"`
	Expire int64  `yaml:"Expire"`
}

func ConfigInit(path string) error {
	var err error
	C, err = ReadYamlConfig(path)
	return err
}

func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err := yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			return nil, err
		}
	}
	return conf, nil
}
