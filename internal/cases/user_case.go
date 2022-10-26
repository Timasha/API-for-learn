package cases

import (
	"context"
)

type UserCase struct {
	userStorage UserStorage
}

func (u *UserCase) New(userStorage UserStorage) {
	u.userStorage = userStorage
}

type User struct {
	Login    string
	Password string
}

func (u *UserCase) CreateUser(ctx context.Context, user User) (string, error) {
	id, err := u.userStorage.CreateUser(ctx, User{
		Login:    user.Login,
		Password: user.Password,
	})

	// if err == storage.ErrUserExist {
	// 	return "", ErrUserExist
	// }

	return id, err
}
func (u *UserCase) ReadUser(ctx context.Context, login string) (User, error) {

	user, err := u.userStorage.ReadUser(ctx, login)

	// if err == storage.ErrUserNotExist {
	// 	return User{}, ErrUserNotExist
	// }

	return User{
		Login:    user.Login,
		Password: user.Password,
	}, err
}
func (u *UserCase) UpdateUser(ctx context.Context, login string, user User) error {
	err := u.userStorage.UpdateUser(ctx, login, User{
		Login:    user.Login,
		Password: user.Password,
	})
	// if err == storage.ErrUserNotExist {
	// 	return ErrUserNotExist
	// }
	return err
}
func (u *UserCase) DeleteUser(ctx context.Context, login string) error {
	err := u.userStorage.DeleteUser(ctx, login)
	// if err == storage.ErrUserNotExist {
	// 	return ErrUserNotExist
	// }
	return err
}
