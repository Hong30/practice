package entities

import (
	"crypto/sha1"
	"encoding/hex"
)

// User 用户
type User struct {
	ID        int64
	UserName  string
	NickName  string
	Password  string
	Salt      string
	CreatedAt int64
}

// GeneratePassword 生成密码
func (u *User) GeneratePassword(originPassword string) string {
	hash := sha1.New()
	hash.Write([]byte(originPassword + u.Salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// Manager 管理员
type Manager struct {
	ID        int64
	UserName  string
	NickName  string
	Password  string
	Salt      string
	CreatedAt int64
}

// GeneratePassword 生成密码
func (m *Manager) GeneratePassword(originPassword string) string {
	hash := sha1.New()
	hash.Write([]byte(originPassword + m.Salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// Record 借书/还书记录
type Record struct {
	ID         int64
	BookID     int64
	BorrowerID int64
	Type       RecordType
	CreatedAt  int64
}
