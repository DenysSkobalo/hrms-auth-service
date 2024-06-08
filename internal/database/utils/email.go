package utils

import "database/sql"

func IsEmailExists(db *sql.DB, email string) (bool, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
