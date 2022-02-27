package user_account

type Portfolio struct {
	ID      int     `json:"id" db:"id"`
	Name    string  `json:"name" db:"name"`
	UserId  int     `json:"userId" db:"user_id"`
	Account float64 `json:"account" db:"account"` //float//decimal in db
}
