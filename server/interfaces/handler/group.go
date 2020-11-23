package handler

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type GroupHandler struct {
	GroupUsecase usecase.GroupUsecase
}

// JSONリクエストを受け取るための構造体一覧
type RequestGroupCreate struct {
	groupUid string `json:"group_uid"`
	name     string `json:"name"`
}

type RequestGroupFindById struct {
	id uint `json:"id"`
}

type RequestGroupChangeName struct {
	id   uint   `json:"id"`
	name string `json:"name"`
}

type RequestGroupAddUser struct {
	groupId uint   `json:"group_id"`
	userUid string `json:"user_uid"`
}

type RequestGroupLeave struct {
	groupId uint   `json:"group_id"`
	userUid string `json:"user_uid"`
}

type RequestGroupDelete struct {
	id uint `json:"id"`
}

func NewGroupHandler(groupUsecase usecase.GroupUsecase) GroupHandler {
	return GroupHandler{GroupUsecase: groupUsecase}
}

func (gh *GroupHandler) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestGroupCreate)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := gh.GroupUsecase.Create(param.name); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (gh *GroupHandler) FindById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestGroupFindById)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		group, err := gh.GroupUsecase.FindById(param.id)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, group)
	}
}

func (gh *GroupHandler) ChangeName() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestGroupChangeName)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		foundSameName, err := gh.GroupUsecase.ChangeName(param.id, param.name)
		if err != nil {
			return err
		}

		if foundSameName == true {
			return ctx.JSON(http.StatusBadRequest, nil)
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (gh *GroupHandler) AddUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestGroupAddUser)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := gh.GroupUsecase.AddUser(param.groupId, param.userUid); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (gh *GroupHandler) Leave() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestGroupLeave)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := gh.GroupUsecase.Leave(param.groupId, param.userUid); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func (gh *GroupHandler) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := new(RequestGroupDelete)

		if err := ctx.Bind(param); err != nil {
			return err
		}

		if err := gh.GroupUsecase.Delete(param.id); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}
