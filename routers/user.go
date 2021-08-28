package routers

import (
  "github.com/gorilla/mux"
  "toonji.com/m/controllers"
)


func SetUsersRoutes(router *mux.Router) *mux.Routers {
  router.HandleFunc("/users/register",controllers.Register).Methods("POST")
  router.HandleFunc("/users/login",controllers.Login).Methods("POST")
  return router
}
