package managers

import (
	"books/domain/entities"
	"database/sql"
)

// FindByUserName 根据用户名查询管理员
func FindByUserName(username string) (*entities.Manager, error) {
	manager := &entities.Manager{}
	row := db.QueryRow("SELECT * FROM `managers` WHERE `username` = ? LIMIT 1", username)
	if err := row.Scan(&manager.ID, &manager.UserName, &manager.NickName, &manager.Password, &manager.Salt, &manager.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return manager, nil
}

// CreateManager 插入管理员j
func CreateManager(manager *entities.Manager) (int64, error) {
	result, err := db.Exec("INSERT INTO `managers`(`username`, `nickname`, `password`, `salt`, `created_at`) VALUES(?, ?, ?, ?, ?)",
		manager.UserName, manager.NickName, manager.Password, manager.Salt, manager.CreatedAt)
	if err != nil {
		return 0, err
	}
	manager.ID, err = result.LastInsertId()
	return manager.ID, err
}
