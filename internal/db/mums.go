package db

import (
	"database/sql"
)

type MumsType string

const (
	Purchase MumsType = "purchase"
	Consumption MumsType = "consumption"
)

const SchemaMums = `
CREATE TABLE IF NOT EXISTS mums (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_account_id INTEGER NOT NULL,
	phaddergrupp_id INTEGER NOT NULL,
	timestamp TEXT NOT NULL DEFAULT (datetime('now')),
	mums_type TEXT NOT NULL,
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id),
    FOREIGN KEY (phaddergrupp_id) REFERENCES phaddergrupps(id)
);`

func CreateMums(db *sql.DB, userAccountID, phaddergruppID int64, mumsType MumsType) (int64, error) {
	res, err := db.Exec(
		`INSERT INTO mums (user_account_id, phaddergrupp_id, mums_type) VALUES (?, ?, ?)`,
		userAccountID,
		phaddergruppID,
		string(mumsType),
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertID()
	return id, err
}

