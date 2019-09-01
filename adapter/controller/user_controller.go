package controller

import (
	"api/adapter/logger"
	"api/adapter/presenter"
	"api/domain"
	"api/usecase/interactor"
)

type UserController struct {
	i interactor.UserInteractor
	l logger.Logger
}

func NewUserController(i interactor.UserInteractor, l logger.Logger) *UserController {
	return &UserController{i, l}
}

func (uc *UserController) SignUp(ctx Context) {
	ep := presenter.NewErrorPresenter(uc.l, ctx)
	u := domain.User{}
	ctx.Bind(&u)
	uc.i.Create(u, ep)
	// login
	ctx.JSON(200, u)
}