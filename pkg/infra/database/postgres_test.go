package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPostgresConnector_String(t *testing.T) {
	want := "host=Host port=5432 user=User password=Password dbname=NameDB sslmode=disable timezone=Timezone"

	conn := &Postgres{
		Host:          "Host",
		Port:          5432,
		User:          "User",
		Password:      "Password",
		NameDB:        "NameDB",
		SSL:           false,
		Timezone:      "Timezone",
		Timeout:       time.Second,
		RetryInterval: time.Second,
	}

	got := conn.String()

	assert.Equal(t, want, got)
}

func TestPostgresConnector_ConnectDB(t *testing.T) {
	t.SkipNow()
}

func TestPostgresConnector_configConn(t *testing.T) {
	t.SkipNow()
}
