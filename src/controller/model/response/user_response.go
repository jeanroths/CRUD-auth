package response 

type UserResponse struct {
	Id string `json:"id" bson:"_id"
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
