package configs

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Application AppConfig      `mapstructure:"application"`
	Server      ServerConfig   `mapstructure:"server"`
	Database    DatabaseConfig `mapstructure:"database"`
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}

	return &config
}
