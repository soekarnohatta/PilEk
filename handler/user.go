package handler

import (
	"PilEk/model"
	"PilEk/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) Login(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	u, err := h.us.GetByUsername(req.User.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newUserLoginResponse(u))
}

func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.us.GetByUsername(userNameFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newUserLoginResponse(u))
}

func (h *Handler) Update(c echo.Context) error {
	req := newUserUpdateRequest()
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	u, err := h.us.GetByUsername(userNameFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	upd := req.populate(u)
	if err := h.us.UpdateUser(upd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) Vote(c echo.Context) error {
	u, err := h.us.GetByUsername(userNameFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	if u.Status == model.SudahMemilih {
		return c.JSON(http.StatusBadRequest, utils.NewError(fmt.Errorf("%v", "you have been electing")))
	}

	idkandidat, _ := strconv.Atoi(c.QueryParam("idkandidat"))
	k, err := h.ks.GetByNomorUrut(idkandidat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if k == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	k.JumlahSuara += 1
	if err := h.ks.UpdateKandidat(k); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	u.Status = model.SudahMemilih
	u.IdKandidat = idkandidat
	if err := h.us.UpdateUser(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.String(http.StatusOK, "ok")
}

func userNameFromToken(c echo.Context) string {
	id, ok := c.Get("user").(string)
	if !ok {
		return ""
	}
	return id
}
