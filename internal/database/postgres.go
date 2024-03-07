// Package database performs database connections and persistence operations.
package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	connMaxIdleTime = time.Minute
	connMaxLifetime = time.Hour
	maxIdleConns    = 100
	maxOpenConns    = 200
)

// Postgres is used on database operations.
type Postgres struct{ *gorm.DB }

// PostgresConnector is the implementation
// that uses the PostgresConnector repository.
type PostgresConnector struct {
	Host          string
	Port          int
	User          string
	Password      string
	NameDB        string
	SSL           bool
	Timezone      string
	Timeout       time.Duration
	RetryInterval time.Duration
}

// String returns the connection string.
func (c *PostgresConnector) String() string {
	var ssl string

	if c.SSL {
		ssl = "enable"
	} else {
		ssl = "disable"
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		c.Host, c.Port, c.User, c.Password, c.NameDB, ssl, c.Timezone,
	)
}

// ConnectDB returns the database connection.
func (c *PostgresConnector) ConnectDB() (*Postgres, error) {
	pg := postgres.Open(c.String())

	gormDB, err := gorm.Open(pg, &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	if err = c.configConn(gormDB); err != nil {
		return nil, err
	}

	return &Postgres{gormDB}, nil
}

// configConn configures the database resolver.
func (c *PostgresConnector) configConn(db *gorm.DB) error {
	sql, err := db.DB()
	if err != nil {
		return err
	}
	sql.SetConnMaxIdleTime(connMaxIdleTime)
	sql.SetConnMaxLifetime(connMaxLifetime)
	sql.SetMaxIdleConns(maxIdleConns)
	sql.SetMaxOpenConns(maxOpenConns)

	return nil
}
