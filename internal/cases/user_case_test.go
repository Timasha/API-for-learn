package cases_test

import (
	"api-for-learn/internal/cases"
	"api-for-learn/internal/storage"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	var mock storage.MockDb = storage.MockDb{
		"Ivan": cases.User{
			Login:    "Ivan",
			Password: "12345",
		},
		"Timofey": cases.User{
			Login:    "Timofey",
			Password: "13524000",
		},
		"User": cases.User{
			Login:    "User",
			Password: "321476129",
		},
	}

	var userCase cases.UserCase
	userCase.New(mock)

	testCases := []struct {
		user  cases.User
		login string
		err   error
	}{
		{
			user: cases.User{
				Login:    "Dmitriy",
				Password: "00112233",
			},
			login: "Dmitriy",
			err:   nil,
		},
		{
			user: cases.User{
				Login:    "Boris",
				Password: "100500",
			},
			login: "Boris",
			err:   nil,
		}, {
			user: cases.User{
				Login:    "Ivan",
				Password: "12345",
			},
			login: "",
			err:   cases.ErrUserExist,
		}, {
			user: cases.User{
				Login:    "Timofey",
				Password: "12345",
			},
			login: "",
			err:   cases.ErrUserExist,
		},
	}
	for _, tC := range testCases {
		login, err := userCase.CreateUser(context.Background(), tC.user)
		assert.Equal(t, tC.login, login)
		assert.Equal(t, tC.err, err)
	}
}

func TestReadUser(t *testing.T) {

	var mock storage.MockDb = storage.MockDb{
		"Ivan": cases.User{
			Login:    "Ivan",
			Password: "12345",
		},
		"Timofey": cases.User{
			Login:    "Timofey",
			Password: "13524000",
		},
		"User": cases.User{
			Login:    "User",
			Password: "321476129",
		},
	}

	var userCase cases.UserCase
	userCase.New(mock)

	testCases := []struct {
		login string
		user  cases.User
		err   error
	}{
		{
			login: "Ivan",
			user: cases.User{
				Login:    "Ivan",
				Password: "12345",
			},
			err: nil,
		},
		{
			login: "Timofey",
			user: cases.User{
				Login:    "Timofey",
				Password: "13524000",
			},
			err: nil,
		},
		{
			login: "1235233",
			user:  cases.User{},
			err:   cases.ErrUserNotExist,
		},
	}
	for _, tC := range testCases {
		user, err := userCase.ReadUser(context.Background(), tC.login)
		assert.Equal(t, tC.user, user)
		assert.Equal(t, tC.err, err)
	}
}

func TestUpdateUser(t *testing.T) {

	var mock storage.MockDb = storage.MockDb{
		"Ivan": cases.User{
			Login:    "Ivan",
			Password: "12345",
		},
		"Timofey": cases.User{
			Login:    "Timofey",
			Password: "13524000",
		},
		"User": cases.User{
			Login:    "User",
			Password: "321476129",
		},
	}

	var userCase cases.UserCase
	userCase.New(mock)

	testCases := []struct {
		login string
		user  cases.User
		err   error
	}{
		{
			login: "Timofey",
			user: cases.User{
				Login:    "Tim",
				Password: "21323135",
			},
			err: nil,
		},
		{
			login: "Ivan",
			user: cases.User{
				Login:    "Ivan",
				Password: "21323135",
			},
			err: nil,
		},
		{
			login: "21354",
			user:  cases.User{},
			err:   cases.ErrUserNotExist,
		},
	}
	for _, tC := range testCases {
		err := userCase.UpdateUser(context.Background(), tC.login, tC.user)
		assert.Equal(t, tC.err, err)
	}
}

func TestDeleteUser(t *testing.T) {

	var mock storage.MockDb = storage.MockDb{
		"Ivan": cases.User{
			Login:    "Ivan",
			Password: "12345",
		},
		"Timofey": cases.User{
			Login:    "Timofey",
			Password: "13524000",
		},
		"User": cases.User{
			Login:    "User",
			Password: "321476129",
		},
	}

	var userCase cases.UserCase
	userCase.New(mock)

	testCases := []struct {
		login string
		err   error
	}{
		{
			login: "Timofey",
			err:   nil,
		},
		{
			login: "Timofey",
			err:   cases.ErrUserNotExist,
		},
		{
			login: "Ivan",
			err:   nil,
		},
	}
	for _, tC := range testCases {
		err := userCase.DeleteUser(context.Background(), tC.login)
		assert.Equal(t, tC.err, err)
	}
}
