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

// Postgres is the implementation
// that uses the Postgres repository.
type Postgres struct {
	Host          string
	Port          int
	User          string
	Password      string
	NameDB        string
	SSL           bool
	Timezone      string
	Timeout       time.Duration
	RetryInterval time.Duration
	*gorm.DB
}

// String returns the connection string.
func (p *Postgres) String() string {
	var ssl string

	if p.SSL {
		ssl = "enable"
	} else {
		ssl = "disable"
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		p.Host, p.Port, p.User, p.Password, p.NameDB, ssl, p.Timezone,
	)
}

// Open returns the database connection.
func (p *Postgres) Open() error {
	pg := postgres.Open(p.String())

	gormDB, err := gorm.Open(pg, &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	if err = p.configConn(gormDB); err != nil {
		return err
	}
	p.DB = gormDB

	return nil
}

// configConn configures the database resolver.
func (p *Postgres) configConn(db *gorm.DB) error {
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
