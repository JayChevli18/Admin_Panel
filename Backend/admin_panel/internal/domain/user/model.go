package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RoleAdmin   Role = "Admin"
	RoleEditior Role = "Editor"
	RoleViewer  Role = "Viewer"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	UserID    int64              `bson:"userId" json:"userId"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Role      Role               `bson:"role" json:"role"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
}

func CollectionName() string {
	return "users"
}
