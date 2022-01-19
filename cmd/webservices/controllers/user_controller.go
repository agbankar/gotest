package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"mocking-goway/internal/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (UserController *UserController) GetBonus(writer http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	fmt.Println(userId)
	writer.Write([]byte("ok"))
	writer.WriteHeader(http.StatusOK)

}
