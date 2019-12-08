package database

import (
	"context"
	"fmt"
	"log"

	"github.com/xyfll7/login/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertAdmin creates a admin.
func (m *MgClient) InsertAdmin(admin *models.Admin) (*models.Admin, error) {
	// Specifies the order in which to return results.
	result, err := m.DB.Collection("admins").InsertOne(
		context.Background(),
		admin,
	)
	fmt.Println(">>>>", result)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// FindAdmin find an admin.
func (m *MgClient) FindAdmin(auth *models.Auth) (*models.Admin, error) {
	var admin models.Admin
	err := m.DB.Collection("admins").FindOne(
		context.Background(),
		bson.M{
			"$or": bson.A{
				bson.M{"name": auth.Adminname},
				bson.M{"email": auth.Adminname},
			},
		},
	).Decode(&admin)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	return &admin, nil
}
