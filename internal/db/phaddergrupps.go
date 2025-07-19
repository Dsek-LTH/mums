package db

import (
	"github.com/Dsek-LTH/mums/internal/config"
)

const SchemaPhaddergrupps = `
CREATE TABLE IF NOT EXISTS phaddergrupps (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	icon_file_path TEXT DEFAULT NULL,
	primary_color TEXT NOT NULL,
	secondary_color TEXT NOT NULL,
	mums_price_n0lla BIGINT NOT NULL,
	mums_price_phadder BIGINT NOT NULL,
	mums_currency TEXT NOT NULL,
	swish_recipient_number TEXT DEFAULT NULL
);`

func (db *DB) CreatePhaddergrupp(exec execer, name string) (int64, error) {
	res, err := exec.Exec(
		`INSERT INTO phaddergrupps (name, primary_color, secondary_color, mums_price_n0lla, mums_price_phadder, mums_currency) VALUES (?, ?, ?, ?, ?, ?)`,
		name,
		config.DefaultPrimaryPhaddergruppColor,
		config.DefaultSecondaryPhaddergruppColor,
		config.DefaultMumsPriceN0lla,
		config.DefaultMumsPricePhadder,
		config.DefaultMumsCurrecy,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()

	db.Emit(DBEvent{
		"phaddergrupps",
		DBCreate,
		nil,
	})

	return id, err
}
