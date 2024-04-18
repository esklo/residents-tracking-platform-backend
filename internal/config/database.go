package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const (
	databaseHostEnvName     = "DB_HOST"
	databasePortEnvName     = "DB_PORT"
	databaseUserEnvName     = "DB_USER"
	databaseDatabaseEnvName = "DB_DATABASE"
	databasePasswordEnvName = "DB_PASSWORD"
)

type DatabaseConfig interface {
	ConnectionString() string
}

type databaseConfig struct {
	host     string
	port     string
	user     string
	database string
	password string
}

func NewDatabaseConfig() (DatabaseConfig, error) {
	host := os.Getenv(databaseHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("database host not found")
	}

	port := os.Getenv(databasePortEnvName)
	if len(port) == 0 {
		return nil, errors.New("database port not found")
	}

	user := os.Getenv(databaseUserEnvName)
	if len(user) == 0 {
		return nil, errors.New("database user not found")
	}

	database := os.Getenv(databaseDatabaseEnvName)
	if len(database) == 0 {
		return nil, errors.New("database database not found")
	}

	password := os.Getenv(databasePasswordEnvName)
	if len(password) == 0 {
		return nil, errors.New("database password not found")
	}

	return &databaseConfig{
		host:     host,
		port:     port,
		user:     user,
		database: database,
		password: password,
	}, nil
}

func (c *databaseConfig) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", c.user, c.password, c.host, c.port, c.database)
}
