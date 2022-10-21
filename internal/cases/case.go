package cases

import (
	"api-for-learn/internal/storage"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserCase struct {
	mongo *storage.Mongo
}

func (u *UserCase) New(mongo *storage.Mongo) {
	u.mongo = mongo
}

type User struct {
	User storage.User
}

func (u *UserCase) CreateUser(ctx context.Context, user User) (string, error) {
	id, err := u.mongo.CreateUser(ctx, user.User)
	if _, ok := err.(mongo.WriteException); ok {
		return "", storage.ErrUserExist
	}
	return id, err
}
func (u *UserCase) ReadUser(ctx context.Context, login string) (User, error) {
	var (
		user User
		err  error
	)
	user.User, err = u.mongo.ReadUser(ctx, login)
	return user, err
}
func (u *UserCase) UpdateUser(ctx context.Context, login string, user User) (int64, error) {
	return u.mongo.UpdateUser(ctx, login, user.User)
}
func (u *UserCase) DeleteUser(ctx context.Context, login string) (int64, error) {
	return u.mongo.DeleteUser(ctx, login)
}
