package usercases

import (
	"books/domain/entities"
	"books/domain/repository/books"
	"database/sql"
	"errors"
	"log"
	"time"
)

// CreateBook 录入图书
func CreateBook(manager *entities.Manager, bookInfo *entities.BookInfo) (*entities.Book, error) {
	// 查找本书是否已经录入过, 如果录入过, 则返回错误提示
	book, err := books.FindByName(bookInfo.Name)
	if err != nil {
		log.Printf("通过书名查找图书失败, name: %s, error: %v", bookInfo.Name, err)
		return nil, err
	}

	if book != nil {
		return nil, errors.New("同名图书已经存在")
	}
	// save book
	book = &entities.Book{
		BookInfo:  *bookInfo,
		Creator:   manager.NickName,
		CreatedAt: time.Now().Unix(),
	}
	book.ID, err = books.Create(book)
	return book, err
}

//ModifyBook 修改图书
func ModifyBook(bookID int64, newName string) error {
	// 为了判断图书id对应的图书数据是否存在， 通过图书id去数据库图书表查询图书
	_, err := books.FindBookByID(bookID)
	if err != nil {
		if err == sql.ErrNoRows {

			return err
		}
		return err
	}
	// book.Status

	// 在数据库里修改图书id对应的图书信息：把书名改成newName
	if err := books.UpdateBookName(bookID, newName); err != nil {
		return err
	}
	return nil
}

// DeleteBook 删除图书
func DeleteBook(manager *entities.Manager, bookID int64) error {
	return nil
}

// ListBookRecords 获取一本图书的借还记录
func ListBookRecords(manager *entities.Manager, bookID int64) ([]*entities.Record, error) {
	return nil, nil
}
