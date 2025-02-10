package repository

func (r *Repository) CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
		title VARCHAR(256) NOT NULL DEFAULT "",
		comment TEXT NOT NULL DEFAULT "",
		repeat VARCHAR(128) NOT NULL DEFAULT "");
		CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);`

	if _, err := r.db.Exec(query); err != nil {
		return err
	}

	return nil
}
