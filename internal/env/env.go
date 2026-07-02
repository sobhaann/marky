package env

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server ServerConfig `toml:"server"`
	Auth   AuthConfig   `toml:"auth"`
}

func NewConfig() *Config {
	cfg := &Config{}
	_, err := toml.DecodeFile("./config/config.toml", cfg)
	if err != nil {
		panic(fmt.Sprintf("NewConfig: failed to decode config.toml file. error: %v", err))
	}
	return cfg
}

type ServerConfig struct {
	Port         int           `toml:"port"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

type AuthConfig struct {
	JWTSecret    string        `toml:"jwt_secret"`
	JWTIssuer    string        `toml:"jwt_issuer"`
	JWTExpiresAt time.Duration `toml:"jwt_expires_at"`
}
