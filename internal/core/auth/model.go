package auth

type User struct {
	ID       int64
	Name     string
	Password string
	IsAdmin  bool
}
