package config

type Config struct {
	Server Server `mapstructure:"server"`
	World  World  `mapstructure:"world"`
	Player Player `mapstructure:"player"`
}

type Server struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

type World struct {
}

type Player struct {
}

func Build(opts ...Option) *Config {
	return Apply(Default(), opts...)
}

func Apply(cfg *Config, opts ...Option) *Config {
	for _, op := range opts {
		cfg = op(cfg)
	}

	return cfg
}

func Default() *Config {
	return &Config{
		Server: Server{
			Address: "0.0.0.0",
			Port:    25565,
		},
	}
}
