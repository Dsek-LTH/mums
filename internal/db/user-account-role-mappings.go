package db

import (
	"database/sql"

	"github.com/Dsek-LTH/mums/internal/roles"
)

const SchemaUserAccountRoleMappings = `
CREATE TABLE IF NOT EXISTS user_account_role_mappings (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_account_id INTEGER NOT NULL,
	user_account_role TEXT NOT NULL,
	FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE
);`

func CreateUserAccountRoleMapping(db *sql.DB, userAccountId int64, userAccountRole models.UserAccountRole) (int64, error) {
	res, err := db.Exec(
		`INSERT INTO user_account_role_mappings (user_account_id, user_account_role) VALUES (?, ?)`,
		userAccountId,
		string(userAccountRole),
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func HasUserAccountRole(db *sql.DB, userId int64, roles []string) (bool, error) {
	rows, err := db.Query(`
		SELECT user_account_role
		FROM user_account_role_mappings
		WHERE user_account_id = ?`, userId)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return false, err
		}
		for _, r := range roles {
			if role == r {
				return true, nil
			}
		}
	}
	return false, nil
}

