package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Db       string
}

func (d *Database) validate() error {
	if d.Host == "" {
		return errors.New("empty db host provided")
	}
	if d.Port == 0 {
		return errors.New("empty db port provided")
	}
	if d.User == "" {
		return errors.New("empty db user provided")
	}
	if d.Password == "" {
		return errors.New("empty db password provided")
	}
	if d.Db == "" {
		return errors.New("empty db name provided")
	}
	return nil
}

type Config struct {
	Database Database
}

func (c *Config) validate() error {
	return multierr.Combine(
		c.Database.validate(),
	)
}

func Parse(filepath string) (*Config, error) {
	// Parse the file
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	// Validate the provided configuration
	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate the config")
	}
	return &cfg, nil
}
