package dto

type RegisterByUsername struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,password"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty" binding:"email"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Password  string `json:"password"`
	Activated bool   `json:"activated"`
}
type LoginByUsername struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
