package db

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	_ "modernc.org/sqlite"
)

var Schemas = []string{
		SchemaCredentials,
		SchemaUsers,
		SchemaGlobalRoleBindings,
		SchemaPhaddergrupper,
		SchemaPhaddergruppBindings,
		SchemaMums,
	}

func InitDB(dbFilePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	_, err = db.Exec(`PRAGMA journal_mode = WAL;`)
	if err != nil {
		return nil, fmt.Errorf("enabling WAL failed: %w", err)
	}
	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		return nil, fmt.Errorf("enabling foreign_keys failed: %w", err)
	}

	for _, schema := range Schemas {
		if _, err := db.Exec(schema); err != nil {
			return nil, fmt.Errorf("failed to create schema: %w", err)
		}
	}

	return db, nil
}

func DBMiddleware(db *sql.DB) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Set("db", db)
            return next(c)
        }
    }
}

func GetDB(c echo.Context) *sql.DB {
	db, ok := c.Get("db").(*sql.DB)
	if !ok || db == nil {
		panic("no database in context")
	}
	return db
}

