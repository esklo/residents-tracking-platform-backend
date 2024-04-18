package provider

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func (s *ServiceProvider) DatabaseConnection() (*sql.DB, error) {
	if s.databaseConnection == nil {
		db, err := sql.Open("postgres", s.DatabaseConfig().ConnectionString())
		if err != nil {
			return nil, err
		}
		s.databaseConnection = db
	}

	return s.databaseConnection, nil
}
