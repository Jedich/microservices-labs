package models

type Model struct {
	ID int `json:"id" gorm:"primarykey"`
}

type User struct {
	Model
	Username     string   `json:"username" gorm:"index"`
	Email        string   `json:"email" gorm:"index"`
	PasswordHash string   `json:"-"`
	IsAdmin      bool     `json:"is_admin"`
	Ratings      []Rating `json:"-" gorm:"foreignKey:UserID"`
}

type Movie struct {
	Model
	Title       string   `json:"title"`
	Genre       string   `json:"genre"`
	ReleaseYear int      `json:"year"`
	Rating      float64  `json:"rating"`
	Ratings     []Rating `json:"-" gorm:"foreignKey:MovieID"`
}

type Rating struct {
	Model
	UserID  int     `json:"user_id"`
	MovieID int     `json:"movie_id"`
	Score   float64 `json:"score"`
}

type Wishlist struct {
	Model
	UserID  int `json:"user_id"`
	MovieID int `json:"movie_id"`
}

type EmailNotification struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}
