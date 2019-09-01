package presenter

import (
	"api/external/framework"
	"api/adapter/viewmodel"
	"api/adapter/logger"
)

type ErrorPresenter struct {
	l logger.Logger
	c framework.Context
}

func NewErrorPresenter(l logger.Logger, c framework.Context) *ErrorPresenter {
	return &ErrorPresenter{l, c}
}

func (e *ErrorPresenter) ViewError(code int, msg string, err error) {
	e.l.Errorf(err.Error())
	e.c.JSON(code, viewmodel.Error{Code: code, Msg: msg})
}

