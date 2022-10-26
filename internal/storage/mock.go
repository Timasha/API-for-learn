package storage

import (
	"api-for-learn/internal/cases"
	"context"
)

type MockDb map[string]interface{}

func (t MockDb) Connect(context.Context, string, string) error {
	return nil
}
func (t MockDb) Close() {}

func (t MockDb) CreateUser(ctx context.Context, user cases.User) (string, error) {
	if _, ok := t[user.Login]; ok {
		return "", cases.ErrUserExist
	}
	t[user.Login] = user
	return user.Login, nil
}
func (t MockDb) ReadUser(ctx context.Context, login string) (cases.User, error) {
	if _, ok := t[login]; !ok {
		return cases.User{}, cases.ErrUserNotExist
	}
	return t[login].(cases.User), nil
}
func (t MockDb) UpdateUser(ctx context.Context, login string, user cases.User) error {
	if _, ok := t[login]; !ok {
		return cases.ErrUserNotExist
	}
	t[login] = user
	return nil
}
func (t MockDb) DeleteUser(ctx context.Context, login string) error {
	if _, ok := t[login]; !ok {
		return cases.ErrUserNotExist
	}
	delete(t, login)
	return nil
}
