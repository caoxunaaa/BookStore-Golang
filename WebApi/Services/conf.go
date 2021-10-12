package Services

import (
	"gopkg.in/yaml.v2"
	"os"
)

var C *Config

type Config struct {
	Jwt JwtConfig `yaml:"jwt"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire int64 `yaml:"expire"`
}

func ConfigInit(path string) error{
	var err error
	C, err = ReadYamlConfig(path)
	return err
}

func ReadYamlConfig(path string)  (*Config,error){
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil,err
	} else {
		err := yaml.NewDecoder(f).Decode(conf)
		if err !=nil{
			return nil, err
		}
	}
	return  conf,nil
}


