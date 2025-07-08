package db

import (
	"database/sql"
	"slices"

	"github.com/Dsek-LTH/mums/internal/roles"
)

const SchemaPhaddergruppMappings = `
CREATE TABLE IF NOT EXISTS phaddergrupp_mappings (
    user_account_id INTEGER NOT NULL,
    phaddergrupp_id INTEGER NOT NULL,
    phaddergrupp_role TEXT NOT NULL,
    mums_available INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (user_account_id, phaddergrupp_id),
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (phaddergrupp_id) REFERENCES phaddergrupps(id) ON DELETE CASCADE
);`

func CreatePhaddergruppMapping(db *sql.DB, userAccountID, phaddergruppID int64, phaddergruppRole roles.PhaddergruppRole) error {
	_, err := db.Exec(
		`INSERT INTO phaddergrupp_mappings (user_account_id, phaddergrupp_id, phaddergrupp_role) VALUES (?, ?, ?)`,
		userAccountID,
		phaddergruppID,
		string(phaddergruppRole),
	)
	return err
}

func ReadPhaddergruppRole(db *sql.DB, userAccountID, phaddergruppID int64) (roles.PhaddergruppRole, error) {
	rows, err := db.Query(`
		SELECT phaddergrupp_role
		FROM phaddergrupp_mappings
		WHERE user_account_id = ? AND phaddergrupp_id = ?`,
		userAccountID, phaddergruppID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", sql.ErrNoRows
	}

	var phaddergruppRole roles.PhaddergruppRole
	if err := rows.Scan(&phaddergruppRole); err != nil {
		return "", err
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return phaddergruppRole, nil
}

