package handler

import (
	"PilEk/model"
	"github.com/labstack/echo/v4"
)

type userUpdateRequest struct {
	User struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Nama       string `json:"nama"`
		IdKelas    int    `json:"idkelas"`
		IdKandidat int    `json:"idkandidat"`
		Status     int    `json:"status"`
		Aktif      bool   `json:"aktif"`
	} `json:"user"`
}

func newUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

func (ru *userUpdateRequest) populate(u *model.User) *model.User {
	u.Username = ru.User.Username
	u.Password = ru.User.Password
	u.Nama = ru.User.Nama
	u.IdKelas = ru.User.IdKelas
	u.IdKandidat = ru.User.IdKandidat
	u.Status = ru.User.Status
	u.Aktif = ru.User.Aktif
	return u
}

func (ru *userUpdateRequest) bind(c echo.Context) error {
	if err := c.Bind(ru); err != nil {
		return err
	}
	if err := c.Validate(ru); err != nil {
		return err
	}
	return nil
}

type userLoginRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
