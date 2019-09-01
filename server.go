package main

import (
	"api/adapter/presenter"
	"api/adapter/repository"
	"api/external/database"
	"api/external/router"
	"api/usecase/interactor"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	l := looger.NewLogger() // 出力先必要
	db := database.NewDB(l)
	up := presenter.NewUserPresenter(l)
	ep := presenter.NewErrorPresenter()
	ur := repository.NewUserRepository(db)
	ui := interactor.NewUserInteractor(up, ur)
	uc := controller.NewUserController(ui, l)
	r := router.NewRouter(uc)
	port, _ := strconv.Atoi(os.Args[1])
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		l.Err
	}
}
