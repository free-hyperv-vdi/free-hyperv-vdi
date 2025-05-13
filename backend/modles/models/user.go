package models

import (
	"encoding/hex"

	"github.com/google/uuid"
)


type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}





func NewUUID() string {
	u, _ := uuid.NewRandom()
	buf := make([]byte, 32)
	hex.Encode(buf[0:8], u[0:4])
	hex.Encode(buf[8:12], u[4:6])
	hex.Encode(buf[12:16], u[6:8])
	hex.Encode(buf[16:20], u[8:10])
	hex.Encode(buf[20:], u[10:])
	return string(buf)
}
