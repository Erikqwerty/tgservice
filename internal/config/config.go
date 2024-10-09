package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config представляет конфигурацию приложения, включая порт, API-ключ для Telegram и идентификатор чата.
type Config struct {
	Port     string `mapstructure:"port"`
	TgApikey string `mapstructure:"tgapikey"`
	ChatID   int64  `mapstructure:"chatid"`
}

// New загружает конфигурацию из файла, расположенного в указанной директории.
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
