package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var mongodbURL string
	var isAuth bool
	if username == "" && password == "" {
		mongodbURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		mongodbURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
		isAuth = true
	}

	clientOptions := options.Client().ApplyURI(mongodbURL)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	// Connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDb due to error: %v", err)
	}

	// Ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDb due to error: %v", err)
	}

	return client.Database(database), nil
}
