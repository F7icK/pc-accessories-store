package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("err with Open DB: %s", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("err with ping DB: %s", err)
	}

	return &Postgres{db: sqlDB}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) Database() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: p.db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("err with Open GORM: %s", err)
	}

	return db, nil
}
