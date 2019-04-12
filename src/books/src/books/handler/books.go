package handler

import (
	"books/domain/entities"
	"books/domain/usercases"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// CreateBook 创建图书的http接口
func CreateBook(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("TOKEN")
	manager, err := loginedManagers.GetManagerByToken(token)
	if err != nil {
		w.Write([]byte("获取管理员登录信息失败:" + err.Error()))
		return
	}

	bookInfo := &entities.BookInfo{}
	if err := json.NewDecoder(r.Body).Decode(bookInfo); err != nil {
		w.Write([]byte("解析参数失败: " + err.Error()))
		return
	}

	log.Printf("开始创建图书, book info: %+v", bookInfo)
	if len(bookInfo.Name) == 0 {
		log.Println("图书名称为空")
		w.Write([]byte("图书名称不能为空"))
		return
	}

	book, err := usercases.CreateBook(manager, bookInfo)
	if err != nil {
		log.Printf("创建图书失败, error: %+v", err)
		w.Write([]byte("创建图书失败: " + err.Error()))
		return
	}
	log.Printf("图书创建成功, book: %+v", book)
	w.Write([]byte(fmt.Sprintf("%+v", book)))
}

// ListBooks 列出图书列表
func ListBooks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	if page == 0 {
		page = 1
	}

	books, err := usercases.ListBooks(int(page), 10)
	if err != nil {
		log.Printf("获取图书列表失败, error: %+v", err)
		w.Write([]byte("获取图书列表失败: " + err.Error()))
		return
	}

	json.NewEncoder(w).Encode(books)
}

// SearchBooksRequest 图书信息
type SearchBooksRequest struct {
	NameOrAuthor string `json:"name_or_author,omitempty"`
	Page         int    `json:"page,omitempty"`
	PerPage      int    `json:"perPage,omitempty"`
}

type SearchBooksResponse struct {
	Code    int              `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Books   []*entities.Book `json:"books,omitempty"`
	Total   int64            `json:"total,omitempty"`
}

// SearchBooks 搜索图书
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("搜索图书")
	//log.Println(r.Body)
	req := &SearchBooksRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		json.NewEncoder(w).Encode(SearchBooksResponse{
			Code:    1,
			Message: "解析参数失败: " + err.Error(),
		})
		return
	}

	if len(req.NameOrAuthor) == 0 {
		json.NewEncoder(w).Encode(SearchBooksResponse{
			Code:    2,
			Message: "参数不能为空",
		})
		return
	}

	log.Printf("req: %+v", req)
	books, total, err := usercases.SearchBooks(req.NameOrAuthor, req.Page, req.PerPage)
	if err != nil {
		json.NewEncoder(w).Encode(SearchBooksResponse{
			Code:    3,
			Message: "搜索图书失败:" + err.Error(),
		})
		return
	}

	log.Println("books:", books)
	json.NewEncoder(w).Encode(SearchBooksResponse{
		Code:    0,
		Message: "suceess",
		Books:   books,
		Total:   total,
	})
}

type ModifyBooksRequest struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ModifyBooksResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// 修改图书
func ModifyBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("搜索图书")
	//log.Println("books:", books)

	// 创建一个修改图书的请求对象，用来接收客户端发上来的请求数据
	req := &ModifyBooksRequest{}

	// 使用json库， 从http request的body里解析json结构到修改图书的请求对象里
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		// 使用json库，向http writer写入用json库编码过后的修改图书的响应信息
		json.NewEncoder(w).Encode(ModifyBooksResponse{
			Code:    1,
			Message: "解析参数失败: " + err.Error(),
		})
		return
	}

	// 判断请求里的书名和书的id是不是空的。如果是空的，就输出一个错误信息
	if len(req.Name) == 0 || req.ID == 0 {
		json.NewEncoder(w).Encode(ModifyBooksResponse{
			Code:    2,
			Message: "参数不能为空",
		})
		return
	}

	// 调用usercases包下的ModifyBook函数去根据图书id来修改书名，改成name值
	err := usercases.ModifyBook(req.ID, req.Name)
	if err != nil {
		json.NewEncoder(w).Encode(ModifyBooksResponse{
			Code:    3,
			Message: "修改图书失败:" + err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(ModifyBooksResponse{
		Code:    0,
		Message: "suceess",
	})
}

type Book struct {
	bookId   int64  `json:"bookId,omitempty"`
	bookName string `json:"bookName,omitempty"`
}
