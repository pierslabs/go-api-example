package domain

type User struct {
	ID    string `json:"id" db:"user_id"`
	Name  string `json:"name" db:"full_name"`
	Email string `json:"email" db:"email"`
}

func NewUser(id, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
