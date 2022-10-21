package api

import (
	"api-for-learn/internal/cases"
	"api-for-learn/internal/storage"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUserApi(ctx context.Context, userCase *cases.UserCase, app *fiber.App) {
	app.Put("/users", CreateUserHandler(ctx, userCase))
	app.Get("/users", ReadUserHandler(ctx, userCase))
	app.Post("/users", UpdateUserHandler(ctx, userCase))
	app.Delete("/users", DeleteUserHandler(ctx, userCase))
}
func CreateUserHandler(ctx context.Context, userCase *cases.UserCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req cases.User
		var resp struct {
			login string `json:"login"`
			err   string `json:"error"`
		}
		var err error
		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)
		resp.login, err = userCase.CreateUser(timeCtx, req)

		if err == storage.ErrUserExist {
			resp.err = err.Error()
			data, _ := json.Marshal(&resp)
			c.Status(400).Write(data)
		} else if err != nil {
			resp.err = "internal server error"
			log.Println(err)
			data, _ := json.Marshal(&resp)
			c.Status(500).Write(data)
		}

		data, _ := json.Marshal(&resp)
		c.Write(data)
		return nil
	}
}
func ReadUserHandler(ctx context.Context, userCase *cases.UserCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			login string `json:"login"`
		}
		var resp struct {
			user cases.User `json:"user"`
			err  string     `json:"error"`
		}
		var err error

		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)

		resp.user, err = userCase.ReadUser(timeCtx, req.login)
		resp.err = err.Error()

		data, _ := json.Marshal(&resp)
		c.Write(data)
		return nil
	}
}
func UpdateUserHandler(ctx context.Context, userCase *cases.UserCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			login string     `json:"login"`
			user  cases.User `json:"user"`
		}
		var resp struct {
			matched int64  `json:"matched"`
			err     string `json:"error"`
		}
		var err error

		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)

		resp.matched, err = userCase.UpdateUser(timeCtx, req.login, req.user)
		resp.err = err.Error()

		data, _ := json.Marshal(&resp)
		c.Write(data)
		return nil
	}
}
func DeleteUserHandler(ctx context.Context, userCase *cases.UserCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			login string `json:"login"`
		}
		var resp struct {
			matched int64  `json:"matched"`
			err     string `json:"error"`
		}
		var err error

		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)

		resp.matched, err = userCase.DeleteUser(timeCtx, req.login)
		resp.err = err.Error()

		data, _ := json.Marshal(&resp)
		c.Write(data)
		return nil
	}
}
