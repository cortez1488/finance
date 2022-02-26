package psqlAuth

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/core/auth"
	"myFinanceTask/internal/handler/rest"
)

type userStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *userStorage {
	return &userStorage{db: db}
}

func (r *userStorage) CreateUser(user rest.UserDTO) (int, error) {

	query := fmt.Sprintf("INSERT INTO %s (name , hashPass) VALUES ($1, $2) RETURNING id", "users")
	row := r.db.QueryRow(query, user.Name, user.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *userStorage) GetUser(username, password string) (auth.User, error) {
	var user auth.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1 AND hashPass = $2", "users")
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *userStorage) IsAdmin(id int64) bool {
	var isAdmin bool
	query := fmt.Sprintf("SELECT isadmin FROM %s WHERE id = $1", "users")
	row := r.db.QueryRow(query, id)
	row.Scan(&isAdmin)

	return isAdmin
}
