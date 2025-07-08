package db

import (
	"database/sql"
)

const SchemaUserProfiles = `
CREATE TABLE IF NOT EXISTS user_profiles (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);`

func CreateUserProfile(db *sql.DB, name string) (int64, error) {
	res, err := db.Exec(`INSERT INTO user_profiles (name) VALUES (?)`,  name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID()
}
