package conf

import (
	"extractor/conf/grafana"
	"extractor/conf/storage"
)

func Init() *Config {
	return &Config{
		Storage: storage.DefaultOptions(),
		Grafana: grafana.DefaultOptions(),
	}
}
