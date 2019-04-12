package usercases

import (
	"books/domain/entities"
	"books/domain/repository/books"
)

// ListBooks 获取图书列表
func ListBooks(page, perPage int) ([]*entities.Book, error) {
	return books.SelectBooks(page, perPage)
}

// SearchBooks 搜索图书
func SearchBooks(nameOrAuthor string, page, perPage int) ([]*entities.Book, int64, error) {
	// 记录搜索热度
	bookList, err := books.SearchBooks(nameOrAuthor, page, perPage)
	if err != nil {
		return nil, 0, err
	}

	total, err := books.CountBooks(nameOrAuthor)
	if err != nil {
		return nil, 0, err
	}

	return bookList, total, nil
}

// handler->usercases->reponsitory->entities 依赖结构
