package usercases

import (
	"books/domain/entities"
	"books/domain/repository/managers"
	"books/domain/repository/users"
	"errors"
	"log"
	"time"
)

// ManagerLogin 管理员登录
func ManagerLogin(username, password string) (manager *entities.Manager, token string, err error) {
	manager, err = managers.FindByUserName(username)
	if err != nil {
		return
	}

	if manager == nil {
		err = errors.New("管理员不存在")
		return
	}

	// 加密过的密码
	if manager.GeneratePassword(password) == manager.Password {
		token = generateToken()
		return
	}

	err = errors.New("密码不正确")
	return
}

// CreateManager 创建管理员
func CreateManager(userName, password, nickName string) (*entities.Manager, string, error) {
	manager := &entities.Manager{
		UserName:  userName,
		NickName:  nickName,
		CreatedAt: time.Now().Unix(),
	}
	manager.Salt = "test"
	manager.Password = manager.GeneratePassword(password)
	if _, err := managers.CreateManager(manager); err != nil {
		log.Println("创建管理员失败:", err)
		return nil, "", err
	}

	return manager, generateToken(), nil
}

func generateToken() string {
	return "tokentest"
}

//用户登录
func UserLogin(username, password string) (user *entities.User, token string, err error) {
	user, err = users.FindByUserName(username)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("用户不存在")
		return
	}

	// 加密过的密码
	if user.GeneratePassword(password) == user.Password {
		token = generateToken()
		return
	}

	err = errors.New("密码不正确")
	return
}

//创建用户
func CreateUser(userName, password, nickName string) (*entities.User, string, error) {
	user := &entities.User{
		UserName:  userName,
		NickName:  nickName,
		Password:  password,
		CreatedAt: time.Now().Unix(),
	}
	user.Salt = "test"
	user.Password = user.GeneratePassword(password)
	if _, err := users.CreateUser(user); err != nil {
		log.Println("创建用户失败:", err)
		return nil, "", err
	}

	return user, generateToken(), nil
}
