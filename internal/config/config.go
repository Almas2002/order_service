package config

import (
	"7kzu-order-service/pkg/constans"
	kafkaClient "7kzu-order-service/pkg/kafka"
	"7kzu-order-service/pkg/postgres"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"fmt"
	"os"
)

var configPath string

type Config struct {
	GRPC        *GRPC               `mapstructure:"grpc"`
	Prometheus  Prometheus          `mapstructure:"prometheus"`
	KafkaTopics *KafkaTopics        `json:"kafkaTopics"`
	Kafka       *kafkaClient.Config `json:"kafka"`
	Postgres    *postgres.Config    `mapstructure:"postgres"`
}

type GRPC struct {
	Port        string `yaml:"port"`
	Development bool   `yaml:"development"`
}
type KafkaTopics struct {
	CreateOrder *kafkaClient.TopicConfig `mapstructure:"createOrder"`
	UpdateOrder *kafkaClient.TopicConfig `mapstructure:"updateOrder"`
}

type Prometheus struct {
	Port string `yaml:"port"`
	Path string `yaml:"path"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constans.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
				return nil, nil
			}
			configPath = fmt.Sprintf("%s/internal/config/config.yml", cwd)
		}

	}

	cfg := &Config{}
	viper.SetConfigType(constans.YAML)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil

}
