package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// New creates a new wrapper for the mongo-go-driver.
func New(connection, dbname string) (*MgClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		return nil, err
	}
	db := client.Database(dbname)
	return &MgClient{DB: db, Client: client, Context: ctx}, nil
}

// MgClient is a wrapper for the mongo-go-driver.
type MgClient struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}

// Close closes the mongo-go-driver connection.
func (d *MgClient) Close() {
	d.Client.Disconnect(d.Context)
}
