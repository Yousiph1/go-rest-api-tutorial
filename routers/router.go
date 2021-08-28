package routers

import (
  "github.com/gorilla/mux"
)


func InitRoutes() *mux.Router {
  router := mux.NewRouter()
  router = routes.SetUsersRoutes(router)
  router = routes.SetTasksRoutes(router)
  router = routes.SetNotesRoutes(router)
  return router
}
