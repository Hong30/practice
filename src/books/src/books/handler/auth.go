package handler

import (
	"books/domain/usercases"
	"encoding/json"
	"log"
	"net/http"
)

// ManagerLoginInfo 管理员登录信息
type ManagerLoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//UserLodinInfo 用户登录信息
type UserLoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// ManagerLogin 管理员登录
func ManagerLogin(w http.ResponseWriter, r *http.Request) {
	loginInfo := &ManagerLoginInfo{}
	if err := json.NewDecoder(r.Body).Decode(loginInfo); err != nil {
		w.Write([]byte(`{"code": 1, "error": "解析参数失败"}`))
		return
	}

	if len(loginInfo.UserName) == 0 || len(loginInfo.Password) == 0 {
		w.Write([]byte(`{"code": 2, "error": "参数不能为空"}`))
		return
	}

	log.Printf("loginInfo: %+v", loginInfo)
	manager, token, err := usercases.ManagerLogin(loginInfo.UserName, loginInfo.Password)
	if err != nil {
		w.Write([]byte("{\"code\": 2, \"error\": \"" + err.Error() + "\"}"))
		return
	}

	loginedManagers.PutLoginedManager(token, manager)
	w.Write([]byte("{\"code\": 0, \"error\": \"\", \"token\": \"" + token + "\"}"))
}

// ManagerRegisterInfo 注册管理员的信息
type ManagerRegisterInfo struct {
	UserName string `json:"user_name,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Password string `json:"password,omitempty"`
}

//UserRegisterInfo 注册用户信息
type UserRegisterInfo struct {
	UserName string `json:"user_name,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Password string `json:"password,omitempty"`
}

// ManagerRegister 注册管理员
func ManagerRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("注册管理员")
	// log.Println(r.Body)
	info := &ManagerRegisterInfo{}
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		w.Write([]byte("解析参数失败: " + err.Error()))
		return
	}

	if len(info.UserName) == 0 || len(info.NickName) == 0 || len(info.Password) == 0 {
		w.Write([]byte("参数不能为空"))
		return
	}

	manager, token, err := usercases.CreateManager(info.UserName, info.Password, info.NickName)
	if err != nil {
		w.Write([]byte("管理员注册失败:" + err.Error()))
		return
	}

	loginedManagers.PutLoginedManager(token, manager)
	log.Println(info.UserName + "注册成功")
	w.Write([]byte("注册并登录成功: " + token))
}

// UserLogin 登录
func UserLogin(w http.ResponseWriter, r *http.Request) {
	loginInfo := &UserLoginInfo{}
	if err := json.NewDecoder(r.Body).Decode(loginInfo); err != nil {
		w.Write([]byte(`{"code": 1, "error": "解析参数失败"}`))
		return
	}

	if len(loginInfo.UserName) == 0 || len(loginInfo.Password) == 0 {
		w.Write([]byte(`{"code": 2, "error": "参数不能为空"}`))
		return
	}

	log.Printf("loginInfo: %+v", loginInfo)
	user, token, err := usercases.UserLogin(loginInfo.UserName, loginInfo.Password)
	if err != nil {
		log.Println(err)
		w.Write([]byte("{\"code\": 2, \"error\": \"" + err.Error() + "\"}"))
		return
	}

	loginedUser.PutLoginedUser(token, user)
	w.Write([]byte("{\"code\": 0, \"error\": \"\", \"token\": \"" + token + "\"}"))
}

// UserRegister 注册用户
func UserRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("注册用户")
	// log.Println(r.Body)
	info := &UserRegisterInfo{}
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		w.Write([]byte("解析参数失败: " + err.Error()))
		return
	}

	if len(info.UserName) == 0 || len(info.NickName) == 0 || len(info.Password) == 0 {
		w.Write([]byte("参数不能为空"))
		return
	}

	user, token, err := usercases.CreateUser(info.UserName, info.Password, info.NickName)
	if err != nil {
		w.Write([]byte("用户注册失败:" + err.Error()))
		return
	}

	loginedUser.PutLoginedUser(token, user)
	log.Println(info.UserName + "注册成功")
	w.Write([]byte("注册并登录成功: " + token))
}

// type ModifyBooksRequest struct {
// 	bookName string `json:"bookName,omitempty"`
// 	bookId   int64  `json:"bookId,omitempty"`
// }
