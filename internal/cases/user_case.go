package cases

import (
	"api-for-learn/internal/storage"
	"context"
	"errors"
)

type UserCase struct {
	repo storage.Repository
}

func (u *UserCase) New(repo storage.Repository) {
	u.repo = repo
}

type User struct {
	Login    string
	Password string
}

var (
	ErrUserExist    error = errors.New("user is exist")
	ErrUserNotExist error = errors.New("user is not exist")
)

func (u *UserCase) CreateUser(ctx context.Context, user User) (string, error) {
	id, err := u.repo.CreateUser(ctx, storage.User{
		Login:    user.Login,
		Password: user.Password,
	})

	if err == storage.ErrUserExist {
		return "", ErrUserExist
	}

	return id, err
}
func (u *UserCase) ReadUser(ctx context.Context, login string) (User, error) {
	var (
		user storage.User
		err  error
	)

	user, err = u.repo.ReadUser(ctx, login)

	if err == storage.ErrUserNotExist {
		return User{}, ErrUserNotExist
	}

	return User{
		Login:    user.Login,
		Password: user.Password,
	}, err
}
func (u *UserCase) UpdateUser(ctx context.Context, login string, user User) error {
	err := u.repo.UpdateUser(ctx, login, storage.User{
		Login:    user.Login,
		Password: user.Password,
	})
	if err == storage.ErrUserNotExist {
		return ErrUserNotExist
	}
	return err
}
func (u *UserCase) DeleteUser(ctx context.Context, login string) error {
	err := u.repo.DeleteUser(ctx, login)
	if err == storage.ErrUserNotExist {
		return ErrUserNotExist
	}
	return err
}
