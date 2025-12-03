package config

import "github.com/spf13/viper"

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}

type DatabaseConfig struct {
	Url string `mapstructure:"url"`
}

type ServerConfig struct {
	BindAddr string `mapstructure:"bind_addr"`
	Debug    bool   `mapstructure:"debug"`
}

func LoadConfig(configPaths ...string) (*Config, error) {
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	// Set defaults as needed
	viper.SetDefault("server.debug", false)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
