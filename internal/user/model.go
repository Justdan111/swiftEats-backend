package user

type User struct {
	ID	   int  `json:"id"`
	Email  string `json:"email"`
	Password string `json:"-"`
	CreatedAt string `json:"created_at"`
}
