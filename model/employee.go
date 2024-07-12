package model

type Employee struct {
	ID        string `json:"id" bson:"id" binding:"required,min=1"`
	FirstName string `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string `json:"last_name" bson:"last_name" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required"`
}
