package config

import (
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Env  EnvConfig
	DB   DBConfig
	Auth AuthConfig
}

type EnvConfig struct {
	AppMode string `mapstructure:"APP_MODE"`
	AppHost string `mapstructure:"APP_HOST"`
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_PORT"`
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type AuthConfig struct {
	ApiKey                string        `mapstructure:"API_KEY_SECRET"`
	AccessTokenSecret     string        `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenDuration   time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenSecret    string        `mapstructure:"REFRESH_TOKEN_SECRET"`
	RefreshTokenDuration  time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	MaxDevice             uint          `mapstructure:"MAX_DEVICE"`
	ResetPasswordDuration time.Duration `mapstructure:"RESET_PASSWORD_DURATION"`
}

func (c *AppConfig) initConfig() (err error) {
	viper.AddConfigPath("./config/env/")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.Unmarshal(&c.Env)
	viper.Unmarshal(&c.DB)
	viper.Unmarshal(&c.Auth)

	return nil
}

func LoadAppConfig() *AppConfig {
	appCfg := AppConfig{}

	err := appCfg.initConfig()
	if err != nil {
		panic(err)
	}

	return &appCfg
}
