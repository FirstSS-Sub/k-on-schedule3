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
	Name string `json:"name"`
}

type RequestUserFindByUid struct {
	Uid string `json:"uid"`
}

type RequestUserGetSchedule struct {
	Uid string `json:"uid"`
}

type RequestUserUpdateSchedule struct {
	Uid string   `json:"uid"`
	Thu []string `json:"thu"`
	Fri []string `json:"fri"`
	Sat []string `json:"sat"`
	Sun []string `json:"sun"`
	Mon []string `json:"mon"`
	Tue []string `json:"tue"`
	Wed []string `json:"wed"`
}

type RequestUserChangeName struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

type RequestUserDelete struct {
	Uid string `json:"uid"`
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

		if err := uh.UserUsecase.Create(param.Name); err != nil {
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

		user, err := uh.UserUsecase.FindByUid(param.Uid)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) GetSchedule() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestUserGetSchedule)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		schedule, err := uh.UserUsecase.GetSchedule(param.Uid)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, schedule)
	}
}

func (uh *UserHandler) UpdateSchedule() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(usecase.RequestUserUpdateSchedule)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := uh.UserUsecase.UpdateSchedule(param); err != nil {
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

		foundSameName, err := uh.UserUsecase.ChangeName(param.Uid, param.Name)
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

		if err := uh.UserUsecase.Delete(param.Uid); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}
