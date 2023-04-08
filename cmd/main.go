package main

import (
	"github.com/imhinotori/the_void/internal/config"
	"github.com/imhinotori/the_void/internal/server"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	k := readConfiguration()

	srv, err := server.CreateServer(config.FromKoanf(k))
	if err != nil {
		log.Fatal().Msgf("failed to create server: %v", err)
	}

	err = srv.AcceptPackets()
	if err != nil {
		log.Error().Msgf("failed to initialize server: %v", err)
	}
}

func readConfiguration() *koanf.Koanf {
	log.Log().Msgf("Reading Configuration...")
	k := koanf.New(".")

	if err := k.Load(file.Provider("./config/server.toml"), toml.Parser()); err != nil {
		log.Fatal().Msgf("failed to read configuration file: %v", err)
	}

	if err := k.Load(file.Provider("./config/world.toml"), toml.Parser()); err != nil {
		log.Fatal().Msgf("failed to read configuration file: %v", err)
	}

	if err := k.Load(file.Provider("./config/player.toml"), toml.Parser()); err != nil {
		log.Fatal().Msgf("failed to read configuration file: %v", err)
	}

	return k
}
