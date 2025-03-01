package models

type User struct {
	Id      int    `json:"id,omitempty"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
}
