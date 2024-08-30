package rest

import (
	"ApiTest/internal/services"
	"net/http"
)

func NewRouter(userService services.UserService, workService services.WorkService) http.Handler {
	handlerUser := NewUserHandlers(userService)
	handlerWork := NewWorkHandlers(workService)

	mux := http.NewServeMux()

	mux.HandleFunc("/createUser", handlerUser.HandleCreateUser)
	mux.HandleFunc("/createWork", handlerWork.HandleCreateWork)

	return mux
}
