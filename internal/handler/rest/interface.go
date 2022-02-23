package rest

type UserService interface {
	CreateUser(UserDTO) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(accessToken string) (int64, error)
}
