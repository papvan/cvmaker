package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"papvan/cvmaker/internal/user"
	"papvan/cvmaker/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	res, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("faild to create user due to error: %v", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	d.logger.Trace(user)
	return "", fmt.Errorf("faild to convert objectid to hex. probably oid: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectid, hex: %s", id)
	}

	filter := bson.M{"_id": oid}
	res := d.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		// TODO: 404
		return u, fmt.Errorf("failed to find one user by id. id: %s due to error: %v", id, err)
	}

	if err := res.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user from db. id: %s due to error: %v", id, err)
	}

	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
