package config

import "log"
import "github.com/caarlos0/env/v6"

type config struct {
	FeaturehubApiKey       string `env:"FEATUREHUB_API_KEY"`
	FeaturehubClientApiKey string `env:"FEATUREHUB_CLIENT_API_KEY"`
	FeaturehubEdgeUrl      string `env:"FEATUREHUB_EDGE_URL"`
	ServicePort            string `env:"SERVICE_PORT" envDefault:"8080"`
	NotificationServiceUrl string `env:"NOTIFICATION_SERVICE_URL"`
}

func LoadConfig() *config {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
