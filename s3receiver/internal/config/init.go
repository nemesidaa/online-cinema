package config

import (
	"os"
	"s3receiver/internal/filestorage"

	"gopkg.in/yaml.v3"
)

var conf *Config

type Config struct {
	*filestorage.FilestorageConfig `yaml:"fs"`
}

func Load(dsn string) error {
	file, err := os.Open(dsn)
	if err != nil {
		return err
	}
	if err := yaml.NewDecoder(file).Decode(conf); err != nil {
		return err
	}
	return nil
}
func Get() *Config {
	if conf == (*Config)(nil) {
		panic("Config isn`t initialized")
	}
	return conf
}
