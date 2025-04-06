package postgres_connect

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
)

type PostgresStorage struct {
	Db *sql.DB
}

func NewPostgresStorage(cfg Postgres, migrationsEnabled bool) (*PostgresStorage, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	storage := &PostgresStorage{Db: db}

	if migrationsEnabled {
		if err := storage.runMigrations(); err != nil {
			return nil, fmt.Errorf("migrations failed: %v", err)
		}
	} else {
		log.Println("Migrations are disabled via flag, skipping.")
	}

	return storage, nil
}

func (s *PostgresStorage) runMigrations() error {

	migrations := &migrate.FileMigrationSource{
		Dir: "/app/postgres_connect/migrations",
	}

	ms := migrate.MigrationSet{
		TableName: "schema_migrations",
	}

	n, err := ms.Exec(s.Db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Printf("Applied %d database migration(s)\n", n)
	return nil
}
