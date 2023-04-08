package config

import (
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

type Option func(cfg *Config) *Config

func FromKoanf(k *koanf.Koanf) Option {
	return func(cfg *Config) *Config {
		err := k.Unmarshal("", cfg)
		if err != nil {
			log.Fatal().Msgf("parsing failed: %v", err)
		}

		return cfg
	}

}
