package books

import (
	"books/domain/entities"
	"database/sql"
	"log"
	"time"
)

// FindByName 通过书名来查找图书
func FindByName(name string) (*entities.Book, error) {
	book := &entities.Book{}
	row := db.QueryRow("SELECT * FROM `books` WHERE `name` = ? LIMIT 1", name)
	if err := row.Scan(&book.ID, &book.Name, &book.Author, &book.Content, &book.Status, &book.Creator, &book.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return book, nil
}

// Create 创建图书
func Create(book *entities.Book) (id int64, err error) {
	book.CreatedAt = time.Now().Unix()
	result, err := db.Exec("INSERT INTO `books`(`name`, `author`, `content`, `status`, `creator`, `created_at`) VALUES(?, ?, ?, ?, ?, ?)",
		book.Name, book.Author, book.Content, book.Status, book.Creator, book.CreatedAt)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	book.ID = id
	return
}

// SelectBooks 查询图书列表
func SelectBooks(page, perPage int) ([]*entities.Book, error) {
	offset := (page - 1) * perPage
	log.Printf("offset: %v, perPage: %v", offset, perPage)
	rows, err := db.Query("SELECT id,name FROM `books` LIMIT ?, ?", offset, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*entities.Book{}
	for rows.Next() {
		book := &entities.Book{}
		if err := rows.Scan(&book.ID, &book.Name); err != nil {
			log.Println("解析图书信息失败: ", err)
			continue
		}
		books = append(books, book)
	}
	log.Printf("查询到图书列表: %+v", books)
	return books, nil
}

func CountBooks(nameOrAuthor string) (int64, error) {
	row := db.QueryRow("SELECT COUNT(*) as total FROM books table WHERE `name` LIKE ? OR author = ?",
		"%"+nameOrAuthor+"%", nameOrAuthor)
	var total int64
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

// SearchBooks 查询图书列表
func SearchBooks(nameOrAuthor string, page, perPage int) ([]*entities.Book, error) {
	offset := (page - 1) * perPage
	log.Printf("offset: %v, perPage: %v", offset, perPage)

	// 分页
	rows, err := db.Query("SELECT * FROM books WHERE `name` LIKE ? OR author = ? LIMIT ?, ?",
		"%"+nameOrAuthor+"%", nameOrAuthor, offset, perPage)
	//rows, err := db.Query("SELECT * FROM books WHERE `name` LIKE ? OR author = ? LIMIT ?, ?",
	//	"%"+nameOrAuthor+"%", nameOrAuthor, offset, perPage)
	//rows, err := db.Query("SELECT id,name FROM `books` WHERE `name` = ? OR author = ? LIMIT ?, ?",
	//	nameOrAuthor, nameOrAuthor, offset, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*entities.Book{}
	for rows.Next() {
		book := &entities.Book{}
		if err := rows.Scan(&book.ID, &book.Name); err != nil {
			log.Println("解析图书信息失败: ", err)
			continue
		}
		books = append(books, book)
	}
	log.Printf("查询到图书列表: %+v", books)
	return books, nil
}

var (
	id   int
	name string
)

func FindBookByID(bookId int64) (*entities.Book, error) {
	row := db.QueryRow("SELECT id,name FROM users WHERE id = ?", id)

	book := &entities.Book{}
	if err := row.Scan(&book.ID, &book.Name); err != nil {
		return nil, err
	}
	log.Printf("图书: %+v", book)
	return book, nil
}

func UpdateBookName(bookId int64, bookName string) error {

	_, err := db.Exec("UPDATE books SET name = ? WHERE id = ?", bookName, bookId)
	if err != nil {
		return err
	}

	//log.Printf("图书更新成功")
	return err
}
