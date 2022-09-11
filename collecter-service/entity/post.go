package entity

type Post struct {
	Meta`json:"meta"`
	Data []Data `json:"data"`
}

type Meta struct {
	Pagination `json:"pagination"`
}

type Pagination struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Links `json:"links"`
}

type Links struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}

type Data struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
