package rest

type UserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdmSymbolDTO struct {
	Abbr     string `json:"abbr"`
	FullName string `json:"fullName"`
}
