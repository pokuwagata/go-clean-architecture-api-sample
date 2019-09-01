package router

import (
	"api/external/framework"
	"api/adapter/controller"
	"net/http"
)

type Router struct {
	userCtl *controller.UserController
}

func (router *Router) NewRouter(userCtl *controller.UserController) {
	http.HandleFunc("/user", router.handleUserRequest)
}

func (router *Router) handleUserRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
		router.userCtl.SignUp(framework.NewContext(w, r))
	}
}