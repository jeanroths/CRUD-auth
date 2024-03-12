package request 

type UserRequest struct {
	Name string `json "name" binding:"required,min=4,max=60"`
	Email string `json "email" binding:""`
	Password string `json "password" binding:""`
}