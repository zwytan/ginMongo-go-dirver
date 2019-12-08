package database

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// CreateUniqueIndex create UniqueIndex
func (d *MgClient) CreateUniqueIndex(collection string, keys ...string) {
	// 复合索引
	keysDoc := bsonx.Doc{}
	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}
	// 创建索引
	idxRet, err := d.DB.Collection(collection).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keysDoc,
			Options: options.Index().SetUnique(true),
		},
		options.CreateIndexes().SetMaxTime(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.Indexes().CreateOne:", idxRet)
}
