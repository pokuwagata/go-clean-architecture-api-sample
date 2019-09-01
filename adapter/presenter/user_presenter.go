package presenter

import (
	"api/adapter/viewmodel"
	"api/adapter/logger"
	"api/domain"
)

type UserPresenter struct {
	l logger.Logger
}


func NewUserPresenter(l logger.Logger) *UserPresenter {
	return &UserPresenter{l}
}

func (p *UserPresenter) ViewUser(u domain.User) *viewmodel.User {
	return &viewmodel.User{
		Id: u.Id,
		Username: u.Username,
	}
}

