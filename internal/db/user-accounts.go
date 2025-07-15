package db

const SchemaUserAccounts = `
CREATE TABLE IF NOT EXISTS user_accounts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_credentials_id INTEGER NOT NULL,
	user_profile_id INTEGER NOT NULL,
	FOREIGN KEY (user_credentials_id) REFERENCES user_credentials(id) ON DELETE CASCADE,
	FOREIGN KEY (user_profile_id) REFERENCES user_profiles(id) ON DELETE CASCADE
);`

func (db *DB) CreateUserAccount(userCredentialsID, userProfileID int64) (int64, error) {
	res, err := db.Exec(
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
