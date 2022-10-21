package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	mongo *mongo.Client
}
type User struct {
	Login    string `bson:"_id" json:"login"`
	Password string `bson:"password" json:"password"`
}

var ErrUserExist error = errors.New("user is exist")

func (m *Mongo) Connect(ctx context.Context, ip, port string) error {
	var connErr error
	opt := options.Client()
	opt.ApplyURI("mongodb://" + ip + ":" + port)
	m.mongo, connErr = mongo.Connect(ctx, opt)
	return connErr
}
func (m *Mongo) Close() {
	m.mongo.Disconnect(context.Background())
}
func (m *Mongo) CreateUser(ctx context.Context, user User) (string, error) {
	res, err := m.mongo.Database("app").Collection("users").InsertOne(ctx, user)
	if res == nil {
		return "", err
	}
	if id, ok := res.InsertedID.(string); ok {
		return id, err
	}
	return "", err
}
func (m *Mongo) ReadUser(ctx context.Context, login string) (User, error) {
	res := m.mongo.Database("app").Collection("users").FindOne(ctx, bson.M{
		"_id": login,
	})
	if res.Err() != nil {
		return User{}, res.Err()
	}
	var resultUser User
	decErr := res.Decode(&resultUser)
	return resultUser, decErr
}
func (m *Mongo) UpdateUser(ctx context.Context, login string, newUser User) (int64, error) {
	res, err := m.mongo.Database("app").Collection("users").ReplaceOne(ctx, bson.M{
		"_id": login,
	}, newUser)
	if res == nil {
		return 0, err
	}
	return res.MatchedCount, err
}
func (m *Mongo) DeleteUser(ctx context.Context, login string) (int64, error) {
	res, err := m.mongo.Database("app").Collection("users").DeleteOne(ctx, bson.M{
		"_id": login,
	})
	if res == nil {
		return 0, err
	}
	return res.DeletedCount, err
}
