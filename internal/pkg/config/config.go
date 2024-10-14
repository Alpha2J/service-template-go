package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	logger "service-template-go/internal/pkg/log"
)

type ServerConfig struct {
	Env   string      `mapstructure:"env"`
	App   AppConfig   `mapstructure:"app"`
	Mysql MySqlConfig `mapstructure:"mysql"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type MySqlConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

const (
	ENV_LOCAL = "local"
	ENV_PROD  = "prod"
	ENV_TEST  = "test"
)

var validEnvs = [...]string{ENV_LOCAL, ENV_PROD, ENV_TEST}

var Config ServerConfig
var env string

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file, %s", err)
		os.Exit(1)
	}

	env = os.Getenv("go_service_env")
	if env == "" {
		env = viper.GetString("env")
	}

	if env == "" {
		env = "local"
	} else {
		isValidEnv := false
		for _, v := range validEnvs {
			if env == v {
				isValidEnv = true
				break
			}
		}
		if !isValidEnv {
			logger.Errorf("Invalid env [%s], only supports one of 'local','prod','test'", env)
			os.Exit(1)
		}
	}

	logger.Infof("Using env: %s", env)

	envConfigFile := fmt.Sprintf("application.%s.yaml", env)
	viper.SetConfigName(envConfigFile)
	viper.AddConfigPath("./configs/")
	if err := viper.MergeInConfig(); err != nil {
		logger.Errorf("Error reading env config file, %s", err)
		os.Exit(1)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		logger.Errorf("Unable to decode into struct, %v", err)
	}
}
