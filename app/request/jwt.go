package request

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	NickName string
	RoleId   uint
	Email    string
}
