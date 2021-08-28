package routers

import (
  "github.com/codegangsta/negroni"
  "github.com/gorila/mux"
  "toonji.com/m/common"
  "toonji.com/m/controllers"
)


func SetTaskRoutes(router *mux.Router) *mux.Router {
  taskRouter = mux.NewRouter()
  taskRouter.HandleFunc("/tasks",controllers.CreateTask).Methods("POST")
  taskRouter.HandleFunc("/tasks",controllers.GetTasks).Methods("GET")
  taskRouter.HandleFunc("/tasks/{id}",controllers.UpdateTask).Methods("PUT")
  taskRouter.HandleFunc("/tasks/{id}",controllers.GetTaskById).Methods("GET")
  taskRouter.HandleFunc("/tasks/users/{id}",controllers.GetTasksByUser).Methods("GET")
  taskRouter.HandleFunc("/tasks/{id}",controllers.DeleteTask).Methods("DELETE")
  router.PathPrefix("/tasks").Handler(negroni.New(
     negroni.HandlerFunc(common.Authorize),
     negroni.Wrap(taskRouter)
    ))

  return router
}
