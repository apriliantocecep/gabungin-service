package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	Db_host     string `mapstructure:"DB_HOST"`
	Db_driver   string `mapstructure:"DB_DRIVER"`
	Db_user     string `mapstructure:"DB_USER"`
	Db_password string `mapstructure:"DB_PASSWORD"`
	Db_name     string `mapstructure:"DB_NAME"`
	Db_port     string `mapstructure:"DB_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
