package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	RabbitMQ struct {
		URL              string   `env:"RABBITMQ_URL" env-default:"amqp://guest:guest@localhost:5672/"`
		Queues           []string `env:"RABBITMQ_QUEUES" env-default:""`
		JournalQueue     string   `env:"RABBITMQ_JOURNAL_QUEUE" env-default:""`
		TransactionQueue string   `env:"RABBITMQ_TRANSACTION_QUEUE" env-default:""`
	}
	MongoDB struct {
		URL                   string `env:"MONGODB_URL" env-default:"mongodb://localhost:27017/"`
		Database              string `env:"MONGODB_DATABASE" env-default:""`
		JournalCollection     string `env:"MONGODB_JOURNAL_COLLECTION" env-default:""`
		TransactionCollection string `env:"MONGODB_TRANSACTION_COLLECTION" env-default:""`
	}
}

var (
	Cfg  *Config
	once sync.Once
)

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		Cfg = &Config{}

		err = cleanenv.ReadConfig(".env", Cfg)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		err = cleanenv.ReadEnv(Cfg)
		if err != nil {
			fmt.Println("Error reading environment variables:", err)
		}
	})
	return Cfg, err
}
