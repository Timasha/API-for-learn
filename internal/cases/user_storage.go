package cases

import (
	"context"
	"errors"
)

type UserStorage interface {
	Connect(context.Context, string, string) error
	Close()
	CreateUser(context.Context, User) (string, error)
	ReadUser(context.Context, string) (User, error)
	UpdateUser(context.Context, string, User) error
	DeleteUser(context.Context, string) error
}

var (
	ErrUserExist    error = errors.New("user is exist")
	ErrUserNotExist error = errors.New("user is not exist")
)
