package models

type User struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    int    `json:"status"`
}
