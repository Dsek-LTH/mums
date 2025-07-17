package db

const SchemaUserProfiles = `
CREATE TABLE IF NOT EXISTS user_profiles (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);`

func (db *DB) CreateUserProfile(exec execer, name string) (int64, error) {
	res, err := exec.Exec(`INSERT INTO user_profiles (name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}

	db.Emit(DBEvent{
		"user_profiles",
		DBCreate,
		nil,
	})

	return res.LastInsertId()
}
