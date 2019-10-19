package app

import "github.com/spf13/viper"

type Config struct {
	configName  string
	configPaths []string

	fresh bool
}

func NewConfig(configName string, configPaths ...string) *Config {
	if configName == "" {
		configName = "config" // default
	}
	return &Config{
		configName:  configName,
		configPaths: configPaths,
	}
}

func (c *Config) init() error {
	if c.fresh {
		return nil
	}
	viper.SetConfigName(c.configName)
	for _, p := range c.configPaths {
		viper.AddConfigPath(p)
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	c.fresh = true
	return nil
}

func (c *Config) Load(key string, obj interface{}) error {
	if err := c.init(); err != nil {
		return err
	}
	return viper.UnmarshalKey(key, &obj)
}
