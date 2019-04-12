package main

import (
	"books/domain/repository/books"
	"books/domain/repository/managers"
	"books/domain/repository/users"
	"books/handler"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	books.Init(db)
	managers.Init(db)
	users.Init(db)
	//创建图书
	http.HandleFunc("/books/create", handler.CreateBook)
	//图书列表
	http.HandleFunc("/books", handler.ListBooks)
	//图书查找
	http.HandleFunc("/books/search", handler.SearchBooks)
	//管理员登陆
	http.HandleFunc("/manager/login", handler.ManagerLogin)
	//注册
	http.HandleFunc("/manager/register", handler.ManagerRegister)
	//修改图书
	http.HandleFunc("/manager/ModifyBook", handler.ModifyBook)
	//用户登录
	http.HandleFunc("/user/login", handler.UserLogin)
	//注册用户
	http.HandleFunc("/user/register", handler.UserRegister)

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Println("启动程序，监听9999端口")

	http.ListenAndServe(":9999", nil)

}
