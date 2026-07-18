package models

type Companies struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	ContactNo string `json:"contactno"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}
