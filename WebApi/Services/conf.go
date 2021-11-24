package Services

import (
	"gopkg.in/yaml.v2"
	"os"
)

var C *Config

type Config struct {
	Host        HostConfig        `yaml:"Host"`
	Jwt         JwtConfig         `yaml:"Jwt"`
	FileStorage FileStorageConfig `yaml:"FileStorage"`
	UserRpc     UserRpcConfig     `yaml:"UserRpc"`
	BookRpc     BookRpcConfig     `yaml:"BookRpc"`
	ActionRpc   ActionRpcConfig   `yaml:"ActionRpc"`
	OrderRpc    OrderRpcConfig    `yaml:"OrderRpc"`
	Redis       []RedisConfig     `yaml:"Redis"` //集群暂时没写，所以只写单点
	Kafka       KafkaConf         `yaml:"Kafka"`
}

type HostConfig struct {
	ListenOn string `yaml:"ListenOn"`
}

type BookRpcConfig struct {
	Host string `yaml:"Host"`
}

type UserRpcConfig struct {
	Host string `yaml:"Host"`
}

type ActionRpcConfig struct {
	Host string `yaml:"Host"`
}

type OrderRpcConfig struct {
	Host string `yaml:"Host"`
}

type JwtConfig struct {
	Secret string `yaml:"Secret"`
	Expire int64  `yaml:"Expire"`
}

type FileStorageConfig struct {
	Path string `yaml:"Path"`
}

type RedisConfig struct {
	Host     string `yaml:"Host"`
	PassWord string `yaml:"PassWord"`
	Type     string `yaml:"Type"`
}

type KafkaConf struct {
	Host  []string `yaml:"Host"`
	Order struct {
		Topic string `yaml:"Topic"`
		Key   string `yaml:"Key"`
	} `yaml:"Order"`
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
