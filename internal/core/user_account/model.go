package user_account

type Portfolio struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	UserId  int    `json:"userId"`
	Account int    `json:"account"`
}
