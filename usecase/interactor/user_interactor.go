package interactor

import (
	"api/domain"
)

type UserRepository interface {
	Store(domain.User) (error)
	// FindById(int) (domain.User, error)
}

type UserPresenter interface {
	ViewUser(domain.User)
}

type ErrorPresenter interface {
	ViewError(int, string)
}

type UserInteractor struct {
	userRepo UserRepository
	userPre UserPresenter
	errPre ErrorPresenter
}

func NewUserInteractor(userPre UserPresenter, userRepo UserRepository, errPre ErrorPresenter) *UserInteractor {
	return &UserInteractor{userRepo, userPre, errPre}
}

func (i *UserInteractor) Create(u domain.User) {
	err := i.userRepo.Store(u)
	if err != nil {
		// TODO: 独自定義のエラーメッセージに変更
		i.errPre.ViewError(500, err.Error())
	}
}