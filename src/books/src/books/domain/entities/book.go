package entities

import "errors"

// BookInfo 待录入的图书信息
type BookInfo struct {
	Name    string `json:"name,omitempty"`
	Author  string `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
}

// BookUpdateSet 图书更新集
type BookUpdateSet struct {
	*BookInfo
	Status BookStatus
}

// Book 已录入的图书
type Book struct {
	BookInfo
	ID        int64
	Status    BookStatus
	Creator   string
	CreatedAt int64
}

// Borrowed 被借出
func (b *Book) Borrowed() error {
	if b.Status == Lended {
		return errors.New("本书已被借走")
	}
	b.Status = Lended
	return nil
}
