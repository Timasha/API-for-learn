package api

import (
	"api-for-learn/internal/cases"
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
		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)
		login, err := userCase.CreateUser(timeCtx, req)

		if err == cases.ErrUserExist {
			resp.err = err.Error()
			data, _ := json.Marshal(&resp)

			c.Status(400).Write(data)
			return nil
		} else if err != nil {
			resp.err = "internal server error"
			log.Println(err)
			data, _ := json.Marshal(&resp)

			c.Status(500).Write(data)
			return nil
		}
		resp.login = login
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

		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)

		user, err := userCase.ReadUser(timeCtx, req.login)

		if err == cases.ErrUserNotExist {
			resp.err = err.Error()
			data, _ := json.Marshal(&resp)

			c.Status(400).Write(data)
			return nil
		} else if err != nil {
			resp.err = "internal server error"
			log.Println(err)
			data, _ := json.Marshal(&resp)

			c.Status(500).Write(data)
			return nil
		}
		resp.user = user

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
			err string `json:"error"`
		}

		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)

		err := userCase.UpdateUser(timeCtx, req.login, req.user)

		if err == cases.ErrUserNotExist {
			resp.err = err.Error()
			data, _ := json.Marshal(&resp)

			c.Status(400).Write(data)
			return nil
		} else if err != nil {
			resp.err = "internal server error"
			log.Println(err)
			data, _ := json.Marshal(&resp)

			c.Status(500).Write(data)
			return nil
		}

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
			err string `json:"error"`
		}

		json.Unmarshal(c.Body(), &req)

		timeCtx, _ := context.WithTimeout(ctx, time.Millisecond*250)

		err := userCase.DeleteUser(timeCtx, req.login)

		if err == cases.ErrUserNotExist {
			resp.err = err.Error()
			data, _ := json.Marshal(&resp)

			c.Status(400).Write(data)
			return nil
		} else if err != nil {
			resp.err = "internal server error"
			log.Println(err)
			data, _ := json.Marshal(&resp)

			c.Status(500).Write(data)
			return nil
		}

		data, _ := json.Marshal(&resp)
		c.Write(data)
		return nil
	}
}
