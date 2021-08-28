package routers

import (
  "github.com/gorilla/mux"
  "github.com/codegangsta/negroni"
  "toonji.com/m/common"
  "toonji.com/m/controllers"
)


func SetNotesRoutes(router *mux.Router) *mux.Router {
   notesRouter = mux.NewRouter
   notesRouter.HandleFunc("/notes",controllers.CreateNote).Methods("POST")
   notesRouter.HandleFunc("/notes",controllers.GetNotes).Methods("GET")
   notesRouter.HandleFunc("/notes/{id}",controllers.GetNoteByID).Methods("GET")
   notesRouter.HandleFunc("/notes/{id}",controllers.UpdateNote).Methods("PUT")
   notesRouter.HandleFunc("/notes/{id}",controllers.DeleteNote).Methods("DELETE")
   notesRouter.HandleFunc("/notes/tasks/{id}",controllers.GetNotesByTask).Methods("GET")
   router.PathPrefix("/notes").Handler(negroni.New(
       negroni.HandlerFunc(common.Authorize),
       negron.Wrap(notesRouter)
     ))
  return router
}
