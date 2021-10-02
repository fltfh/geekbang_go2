package config

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

var Config appConfig

const (
	defaultServerPort = 8080
	defaultEchoSQL = false

)

type AppInfo struct {
	Name string
	Version string
	Host string
	Port string
}

type appConfig struct {
	ServerPort int `mapstructure:"server_port"`
	App        AppInfo
	DbDsn      string    `mapstructure:"db_dsn"`
	EchoSQL    bool   `mapstructure:"echo_sql"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DbDsn, validation.Required),
	)
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()

	v.SetConfigName("app")
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	setDefaultConfig(v)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	err := Config.Validate()
	if err != nil {
		return err
	}
	return nil
}

func setDefaultConfig(v *viper.Viper) {
	v.SetDefault("server_port", defaultServerPort)
	v.SetDefault("echo_sql", defaultEchoSQL)
}
