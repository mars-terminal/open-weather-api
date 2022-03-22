package config

import "github.com/spf13/viper"

func NewConfig(configType, configPath, configName string) error {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
