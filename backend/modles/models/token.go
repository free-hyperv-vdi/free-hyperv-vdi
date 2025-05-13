package models


type Token struct {
	ID           string `json:"id"`
	Value        string `json:"value"`
	Cryptoperiod int64  `json:"cryptoperiod"`
	CreatedTime  int64  `json:"created_time"`
	UserId       string `json:"user_id"`
}
