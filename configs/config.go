package configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

// Config for configurations
type Config struct {
	Name string `yaml:"name"`
	Environment string
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Database Database `yaml:"database"`
	//Asterisk asterisk.Config `yaml:"asterisk"`
}

// Init is using to initialize the configs
func loadConfigFile(env string) (*Config, error) {
	file := "configs/config.yaml"
	var appConfig *Config

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var options map[string]Config
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return nil, err
	}
	opt := options[env]
	opt.Environment = env
	appConfig = &opt

	return appConfig, nil
}

func NewConfig() (*Config, error)  {
	return loadConfigFile("dev")
}