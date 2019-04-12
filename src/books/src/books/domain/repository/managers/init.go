package managers

import "database/sql"

var db *sql.DB

// Init 初始化包
func Init(_db *sql.DB) {
	db = _db
}
