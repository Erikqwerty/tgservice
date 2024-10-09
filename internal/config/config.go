package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"port"`
	TgApikey string `mapstructure:"tgapikey"`
	ChatID   int64  `mapstructure:"chatid"`
}

func New(pathdir string) (*Config, error) {
	conf := &Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(pathdir)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения файла конфигурации: %v", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Ошибка заполнения структуры конфигурации: %v", err)
	}

	return conf, nil
}
