package entity

type Book struct {
	Title    string  `json:"title"`
	Author   *string `json:"author"`
	Overview *string `json:"overview"`
}
