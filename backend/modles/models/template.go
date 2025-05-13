package models


type Template struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	UserName string `json:"user_name"`
	UserPwd  string `json:"user_pwd"`
}
