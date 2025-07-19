package db

const SchemaUserAccounts = `
CREATE TABLE IF NOT EXISTS user_accounts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_credentials_id INTEGER NOT NULL UNIQUE,
	user_profile_id INTEGER NOT NULL UNIQUE,
	FOREIGN KEY (user_credentials_id) REFERENCES user_credentials(id) ON DELETE CASCADE,
	FOREIGN KEY (user_profile_id) REFERENCES user_profiles(id) ON DELETE CASCADE
);`

func (db *DB) CreateUserAccount(exec execer, userCredentialsID, userProfileID int64) (int64, error) {
	res, err := exec.Exec(
		`INSERT INTO user_accounts (user_credentials_id, user_profile_id) VALUES (?, ?)`,
		userCredentialsID,
		userProfileID,
	)
	if err != nil {
		return 0, err
	}

	db.Emit(DBEvent{
		"user_accounts",
		DBCreate,
		nil,
	})

	return res.LastInsertId()
}

func (db *DB) ReadUserAccountIDByUserCredentialsID(q queryer, userCredentialsID int64) (int64, error) {
	row := q.QueryRow(
		`SELECT id FROM user_accounts WHERE user_credentials_id = ?`,
		userCredentialsID,
	)

	var userAccountID int64
	if err := row.Scan(&userAccountID); err != nil {
		return 0, err
	}

	db.Emit(DBEvent{
		"user_accounts",
		DBRead,
		userAccountID,
	})

	return userAccountID, nil
}

func (db *DB) ReadUserProfileNameByUserAccountID(q queryer, userAccountID int64) (string, error) {
	row := q.QueryRow(`
		SELECT p.name
		FROM user_profiles AS p
		JOIN user_accounts AS a ON p.id = a.user_profile_id
		WHERE a.id = ?`,
		userAccountID,
	)

	var name string
	if err := row.Scan(&name); err != nil {
		return "", err
	}

	db.Emit(DBEvent{
		"user_profiles",
		DBRead,
		nil,
	})

	return name, nil
}
