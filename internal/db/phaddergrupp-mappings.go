package db

import (
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

func (db *DB) CreatePhaddergruppMapping(exec execer, userAccountID, phaddergruppID int64, phaddergruppRole roles.PhaddergruppRole) error {
	_, err := exec.Exec(
		`INSERT INTO phaddergrupp_mappings (user_account_id, phaddergrupp_id, phaddergrupp_role) VALUES (?, ?, ?)`,
		userAccountID,
		phaddergruppID,
		string(phaddergruppRole),
	)

	db.Emit(DBEvent{
		"phaddergrupp_mappings",
		DBCreate,
		nil,
	})

	return err
}

func (db *DB) ReadPhaddergruppRole(q queryer, userAccountID, phaddergruppID int64) (roles.PhaddergruppRole, error) {
	row := q.QueryRow(`
		SELECT phaddergrupp_role
		FROM phaddergrupp_mappings
		WHERE user_account_id = ? AND phaddergrupp_id = ?`,
		userAccountID, phaddergruppID)

	var phaddergruppRole roles.PhaddergruppRole
	if err := row.Scan(&phaddergruppRole); err != nil {
		return "", err
	}

	db.Emit(DBEvent{
		"phaddergrupp_mappings",
		DBRead,
		nil,
	})

	return phaddergruppRole, nil
}
