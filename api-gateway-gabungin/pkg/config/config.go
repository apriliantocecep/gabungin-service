package config

import "github.com/spf13/viper"

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthSerivceUrl    string `mapstructure:"AUTH_SERVICE_URL"`
	FamilyServiceurl  string `mapstructure:"FAMILY_SERVICE_URL"`
	AccountServiceurl string `mapstructure:"ACCOUNT_SERVICE_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
