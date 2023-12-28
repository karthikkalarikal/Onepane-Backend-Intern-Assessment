package models

type Comments struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// for users

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street string `json:"street"`
		Suite  string `json:"suite"`
		City   string `json:"city"`
		Zip    string `json:"zipcode"`
		Geo    Geo    `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

// Respose struct

type Response struct {
	PostId        int    `json:"postId"`
	PostName      string `json:"postName"`
	CommentsCount int    `json:"commentsCount"`
	UserName      string `json:"userName"`
	Body          string `json:"body"`
}
