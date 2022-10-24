package storage

import "context"

type MockDb map[string]interface{}

func (t MockDb) Connect(context.Context, string, string) error {
	return nil
}
func (t MockDb) Close() {}

func (t MockDb) CreateUser(ctx context.Context, user User) (string, error) {
	if _, ok := t[user.Login]; ok {
		return "", ErrUserExist
	}
	t[user.Login] = user
	return user.Login, nil
}
func (t MockDb) ReadUser(ctx context.Context, login string) (User, error) {
	if _, ok := t[login]; !ok {
		return User{}, ErrUserNotExist
	}
	return t[login].(User), nil
}
func (t MockDb) UpdateUser(ctx context.Context, login string, user User) error {
	if _, ok := t[login]; !ok {
		return ErrUserNotExist
	}
	t[login] = user
	return nil
}
func (t MockDb) DeleteUser(ctx context.Context, login string) error {
	if _, ok := t[login]; !ok {
		return ErrUserNotExist
	}
	delete(t, login)
	return nil
}
