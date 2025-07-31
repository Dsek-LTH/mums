package db

type MumsType string

const (
	Purchase    MumsType = "purchase"
	Consumption MumsType = "consumption"
)

const SchemaMums = `
CREATE TABLE IF NOT EXISTS mums (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	user_account_id INTEGER NOT NULL,
	phaddergrupp_id INTEGER NOT NULL,
	mums_quantity INTEGER NOT NULL,
	mums_type TEXT NOT NULL,
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id),
    FOREIGN KEY (phaddergrupp_id) REFERENCES phaddergrupps(id)
);`

func (db *DB) CreateMums(exec execer, userAccountID, phaddergruppID, mumsQuantity int64, mumsType MumsType) (int64, error) {
	res, err := exec.Exec(
		`INSERT INTO mums (user_account_id, phaddergrupp_id, mums_quantity, mums_type) VALUES (?, ?, ?, ?)`,
		userAccountID,
		phaddergruppID,
		mumsQuantity,
		mumsType,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()

	db.Emit(DBEvent{
		"mums",
		DBCreate,
		nil,
	})

	return id, err
}
