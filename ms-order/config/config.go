package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type dbcfg struct {
	DSN         string `mapstructure:"DB_DSN"`
	DBName      string `mapstructure:"DB_NAME"`
	MinPoolSize int    `mapstructure:"DB_MIN_POOL_SIZE"`
	MaxPoolSize int    `mapstructure:"DB_MAX_POOL_SIZE"`
	MaxIdleTime string `mapstructure:"DB_MAX_IDLE_TIME"`
}

type xenditcfg struct {
	ApiKey string `mapstructure:"XENDIT_API_KEY"`
}

type Config struct {
	Port        int       `mapstructure:"PORT"`
	Environment string    `mapstructure:"ENVIRONMENT"`
	Db          dbcfg     `mapstructure:",squash"`
	Xendit      xenditcfg `mapstructure:",squash"`
}

func New() (Config, error) {
	cfg := Config{}

	viper.SetDefault("PORT", os.Getenv("PORT"))
	viper.SetDefault("ENVIRONMENT", "development")

	viper.SetDefault("DB_MIN_POOL_SIZE", 25)
	viper.SetDefault("DB_MAX_POOL_SIZE", 25)
	viper.SetDefault("DB_MAX_IDLE_TIME", "15m")

	viper.SetDefault("DB_DSN", os.Getenv("DB_DSN"))
	viper.SetDefault("DB_NAME", os.Getenv("DB_NAME"))

	viper.SetDefault("XENDIT_API_KEY", os.Getenv("XENDIT_API_KEY"))

	viper.AutomaticEnv()

	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	if err := checkUnset(); err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.IgnoreUntaggedFields = true
		dc.ErrorUnset = true
	}); err != nil {
		return Config{}, err
	}

	if structs.HasZero(&cfg) {
		return Config{}, fmt.Errorf("config type has zero value")
	}
	return cfg, nil
}

func checkUnset() error {
	var listunset = []string{}
	for key, val := range viper.AllSettings() {
		if val == "" {
			listunset = append(listunset, strings.ToUpper(key))
		}
	}

	if len(listunset) != 0 {
		envs := pie.Join(listunset, " | ")
		return fmt.Errorf("ENVIRONMENT NOT SET: %v", envs)
	}
	return nil
}
