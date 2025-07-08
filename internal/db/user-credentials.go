package db

import (
	"database/sql"
)

const SchemaUserCredentials = `
CREATE TABLE IF NOT EXISTS user_credentials (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	hashword TEXT NOT NULL
);`

func CreateUserCredentials(db *sql.DB, email string, hashword string) (int64, error) {
	res, err := db.Exec(
		`INSERT INTO user_credentials (email, hashword) VALUES (?, ?)`,
		email,
		hashword,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}
