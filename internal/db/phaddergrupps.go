package db

import (
	"database/sql"

	"github.com/Dsek-LTH/mums/internal/config"
)

const SchemaPhaddergrupps = `
CREATE TABLE IF NOT EXISTS phaddergrupps (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	icon_file_path TEXT DEFAULT NULL,
	primary_color TEXT NOT NULL,
	secondary_color TEXT NOT NULL,
	mums_price BIGINT NOT NULL,
	mums_currency TEXT NOT NULL,
	swish_recipient_number TEXT DEFAULT NULL,
	swish_recipient_name TEXT DEFAULT NULL
);`

func NewPhaddergrupp(db *sql.DB, name string) (int64, error) {
	res, err := db.Exec(
		`INSERT INTO phaddergrupps (name, primary_color, secondary_color, mums_price, mums_currency) VALUES (?, ?, ?, ?, ?)`,
		name,
		config.DefaultPrimaryPhaddergruppColor,
		config.DefaultSecondaryPhaddergruppColor,
		config.DefaultMumsPrice,
		config.DefaultMumsCurrecy,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

