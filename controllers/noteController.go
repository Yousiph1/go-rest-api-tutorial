package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"

  "toonji.com/m/common"
  "toonji.com/m/models"
)


func CreateNote(w http.ResponseWriter, r *http.Request) {

    note := NoteResource{}
    err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
      common.WriteError(w,err,"something went wrong",http.StatusBadRequest)
      return
    }
    ctx := NewContext()
    defer ctx.Close()
    c := ctx.DBCollection("notes")
    repo := data.NoteRepository{&c}
    taskNote := models.TaskNote{}
    taskNote.TaskID = note.Data.TaskID
    taskNote.Description = note.Data.Description
    err = repo.Create(&taskNote)
    j, err := json.Marshall(note)
    if err != nil {
      common.WriteError(w,err,"something went wrong",500)
      return
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(j)
}


func GetNotes(w http.ResponseWriter, r *http.Request) {
  ctx := NewContext()
  defer ctx.Close()
  c := ctx.DBCollection("notes")
  repo := data.NoteRepository{&c}
  notes := repo.GetAllNotes()
  j, err := json.Marshall(&NotesResource{Data:notes})
  if err != nil {
    common.WriteError(w,err,"something went wrong",500)
    return
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusOk)
  w.Write(j)
}


func GetNoteByID(w http.ResponseWriter, r *http.Request) {
   vars := mux.Var(r)
   id := vars["id"]
   ctx := NewContext()
   defer ctx.Close()
   c := ctx.DBCollection("notes")
   repo := data.NoteRepository{&c}
   note, err := GetNote(id)
   taskNote = NoteModel{}
   taskNote.TaskID = note.TaskID
   taskNote.Description = note.Description
   j, err := json.Marshall(&NoteResource{Data:taskNote})
   if err != nil {
     common.WriteError(w,err,"something went wrong",500)
     return
   }
   w.Header().Set("Content-Type","application/json")
   w.WriteHeader(http.StatusOk)
   w.Write(j)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
  vars := mux.Var(r)
  id := vars["id"]
  note := models.TaskNote{}
  err := json.NewDecoder(r.Body).Decode(&note)
  if err != nil {
    common.WriteError(w,err,"something went wrong",400)
    return
  }
  ctx := NewContext()
  defer ctx.Close()
  c := ctx.DBCollection("notes")
  repo := data.NoteRepository{&c}
  err := repo.Update(id,note)
  if err != nil {
    common.WriteError(w,err,"something went wrong",400)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
  vars := mux.Var(r)
  id := vars["id"]
  ctx := NewContext()
  defer ctx.Close()
  c := ctx.DBCollection("notes")
  repo := data.NoteRepository{&c}
  err := repo.Delete(id)
  if err != nil {
    common.WriteError(w,err,"something went wrong",400)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}


func GetNotesByTask(w http.ResponseWriter, r *http.Request) {
  vars := mux.Var(r)
  id := vars["id"]
  ctx := NewContext()
  defer ctx.Close()
  c := ctx.DBCollection("notes")
  repo := data.NoteRepository{&c}
  notes := repo.GetNotesByTask(id)
  j, err := json.Marshall(&NotesResource{Data:notes})
  if err != nil {
    common.WriteError(w,err,"something went wrong",500)
    return
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusOk)
  w.Write(j)
}
