package config

import "github.com/caarlos0/env/v6"

// INFO:環境変数
type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

// INFO:コンストラクタ
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
