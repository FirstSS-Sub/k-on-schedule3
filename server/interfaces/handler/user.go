package handler

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

// JSONリクエストを受け取るための構造体一覧
type RequestCreate struct {
	name string `json:"name"`
}

type RequestFindById struct {
	id uint `json:"id"`
}

type RequestUpdateSchedule struct {
	id  uint   `json:"id"`
	thu string `json:"thu"`
	fri string `json:"fri"`
	sat string `json:"sat"`
	sun string `json:"sun"`
	mon string `json:"mon"`
	tue string `json:"tue"`
	wed string `json:"wed"`
}

type RequestChangeName struct {
	id   uint   `json:"id"`
	name string `json:"name"`
}

type RequestDelete struct {
	id uint `json:"id"`
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{UserUsecase: userUsecase}
}

func (uh *UserHandler) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestCreate)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.Create(param.name); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (uh *UserHandler) FindById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestFindById)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		user, err := uh.UserUsecase.FindById(param.id)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) UpdateSchedule() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestUpdateSchedule)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.UpdateSchedule(param.id, param.thu, param.fri, param.sat, param.sun, param.mon, param.tue, param.wed); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (uh *UserHandler) ChangeName() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestChangeName)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		foundSameName, err := uh.UserUsecase.ChangeName(param.id, param.name)
		if err != nil {
			return err
		}

		if foundSameName == true {
			return ctx.JSON(http.StatusBadRequest, nil)
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (uh *UserHandler) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestDelete)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.Delete(param.id); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}
