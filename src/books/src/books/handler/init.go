package handler

import (
	"books/domain/entities"
	"errors"
	"sync"
)

var loginedManagers *LoginedManagers
var loginedUser *LoginedUser

func init() {
	loginedManagers = NewLoginedManagers()
}

func init() {
	loginedUser = NewLoginedUser()
}

type LoginedManagers struct {
	locker   sync.RWMutex
	managers map[string]*entities.Manager
}

type LoginedUser struct {
	locker sync.RWMutex
	user   map[string]*entities.User
}

func NewLoginedManagers() *LoginedManagers {
	return &LoginedManagers{
		managers: make(map[string]*entities.Manager),
	}
}

func NewLoginedUser() *LoginedUser {
	return &LoginedUser{
		user: make(map[string]*entities.User),
	}
}

func (m *LoginedManagers) PutLoginedManager(token string, manager *entities.Manager) {
	m.locker.Lock()
	m.managers[token] = manager
	m.locker.Unlock()
}

func (m *LoginedUser) PutLoginedUser(token string, user *entities.User) {
	m.locker.Lock()
	m.user[token] = user
	m.locker.Unlock()
}

// GetManagerByToken 获取登录后的管理员信息

func (m *LoginedManagers) GetManagerByToken(token string) (*entities.Manager, error) {
	if len(token) == 0 {
		return nil, errors.New("token为空")
	}
	// 读写锁
	m.locker.RLock()
	manager, ok := m.managers[token]
	m.locker.RUnlock()
	if !ok {
		return nil, errors.New("管理员未登录")
	}
	return manager, nil
}

// GetUSerByToken 获取登录后的信息
func (m *LoginedUser) GetUserByToken(token string) (*entities.User, error) {
	if len(token) == 0 {
		return nil, errors.New("token为空")
	}
	// 读写锁
	m.locker.RLock()
	user, ok := m.user[token]
	m.locker.RUnlock()
	if !ok {
		return nil, errors.New("用户未登录")
	}
	return user, nil
}
