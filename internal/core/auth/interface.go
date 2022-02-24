package auth

import "myFinanceTask/internal/handler/rest"

type Authorization interface {
	CreateUser(user rest.UserDTO) (int, error)
	GetUser(username, password string) (User, error)
	IsAdmin(id int64) bool
}
