package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Env  EnvConfig
	DB   DBConfig
	Auth AuthConfig
}

type EnvConfig struct {
	AppMode   string `envconfig:"APP_MODE"`
	AppHost   string `envconfig:"APP_HOST"`
	AppName   string `envconfig:"APP_NAME"`
	AppPort   string `envconfig:"APP_PORT"`
	SecretKey string `envconfig:"SECRET_KEY"`
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Name     string `envconfig:"DB_NAME"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
}

type AuthConfig struct {
	ApiKey                  string        `envconfig:"API_KEY_SECRET"`
	AccessTokenSecret       string        `envconfig:"ACCESS_TOKEN_SECRET"`
	AccessTokenDuration     time.Duration `envconfig:"ACCESS_TOKEN_DURATION"`
	RefreshTokenSecret      string        `envconfig:"REFRESH_TOKEN_SECRET"`
	RefreshTokenDuration    time.Duration `envconfig:"REFRESH_TOKEN_DURATION"`
	MaxDevice               uint          `envconfig:"MAX_DEVICE"`
	RESET_PASSWORD_DURATION time.Duration `envconfig:"RESET_PASSWORD_DURATION"`
}

func (c *AppConfig) Init() {
	envconfig.MustProcess("", &c.Env)
	envconfig.MustProcess("", &c.DB)
	envconfig.MustProcess("", &c.Auth)
}

func LoadAppConfig() *AppConfig {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/../../../config/env/.env", basePath))
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	appCfg := AppConfig{}
	appCfg.Init()
	return &appCfg
}
