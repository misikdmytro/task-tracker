package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/misikdmytro/task-tracker/internal/config"
)

type ConnectionFactory interface {
	NewDB() (*sqlx.DB, error)
}

type connectionFactory struct {
	config config.DatabseConfig
}

func NewConnectionFactory(c config.DatabseConfig) ConnectionFactory {
	return &connectionFactory{config: c}
}

func (f *connectionFactory) NewDB() (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		f.config.Host,
		f.config.Port,
		f.config.User,
		f.config.Password,
		f.config.Db,
		f.config.SSL,
	)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
