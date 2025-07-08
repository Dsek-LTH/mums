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

func CreateUserAccountRoleMapping(db *sql.DB, userAccountID int64, userAccountRole roles.UserAccountRole) (int64, error) {
	res, err := db.Exec(
		`INSERT INTO user_account_role_mappings (user_account_id, user_account_role) VALUES (?, ?)`,
		userAccountID,
		string(userAccountRole),
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func ReadUserAccountRoles(db *sql.DB, userAccountID int64) ([]roles.UserAccountRole, error) {
	rows, err := db.Query(`
		SELECT user_account_role
		FROM user_account_role_mappings
		WHERE user_account_id = ?`, userAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userAccountRoles []roles.UserAccountRole

	for rows.Next() {
		var userAccountRole roles.UserAccountRole
		if err := rows.Scan(&userAccountRole); err != nil {
			return nil, err
		}
		userAccountRoles = append(userAccountRoles, userAccountRole)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userAccountRoles, nil
}
