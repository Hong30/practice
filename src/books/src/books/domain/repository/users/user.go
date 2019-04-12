package users

import (
	"books/domain/entities"
	"database/sql"
)

func FindByUserName(username string) (*entities.User, error) {
	user := &entities.User{}
	row := db.QueryRow("SELECT * FROM `user` WHERE `username` = ? LIMIT 1", username)
	if err := row.Scan(&user.ID, &user.UserName, &user.NickName, &user.Password, &user.Salt, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func CreateUser(user *entities.User) (int64, error) {
	result, err := db.Exec("INSERT INTO `users`(`username`, `nickname`, `password`, `salt`, `created_at`) VALUES(?, ?, ?, ?, ?)",
		user.UserName, user.NickName, user.Password, user.Salt, user.CreatedAt)
	if err != nil {
		return 0, err
	}
	user.ID, err = result.LastInsertId()
	return user.ID, err
}
