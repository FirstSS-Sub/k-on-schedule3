package handler

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

// JSONリクエストを受け取るための構造体一覧
type RequestUserCreate struct {
	name string `json:"name"`
}

type RequestUserFindByUid struct {
	uid string `json:"uid"`
}

type RequestUserUpdateSchedule struct {
	uid string `json:"uid"`
	thu string `json:"thu"`
	fri string `json:"fri"`
	sat string `json:"sat"`
	sun string `json:"sun"`
	mon string `json:"mon"`
	tue string `json:"tue"`
	wed string `json:"wed"`
}

type RequestUserChangeName struct {
	uid  string `json:"uid"`
	name string `json:"name"`
}

type RequestUserDelete struct {
	uid string `json:"uid"`
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{UserUsecase: userUsecase}
}

func (uh *UserHandler) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestUserCreate)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.Create(param.name); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (uh *UserHandler) FindByUid() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestUserFindByUid)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		user, err := uh.UserUsecase.FindByUid(param.uid)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) UpdateSchedule() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestUserUpdateSchedule)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.UpdateSchedule(param.uid, param.thu, param.fri, param.sat, param.sun, param.mon, param.tue, param.wed); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (uh *UserHandler) ChangeName() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestUserChangeName)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		foundSameName, err := uh.UserUsecase.ChangeName(param.uid, param.name)
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
		param := new(RequestUserDelete)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.Delete(param.uid); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}
