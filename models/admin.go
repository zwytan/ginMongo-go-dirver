package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Admin demo
type Admin struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Roles    []string           `bson:"roles" json:"roles"`
}

// New is an instance
func (u *Admin) New() *Admin {
	return &Admin{
		ID:       primitive.NewObjectID(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Roles:    u.Roles,
	}
}
