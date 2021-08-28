package controllers

import (
  "encoding/json"
  "log"
  "net/http"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/gorilla/mux"

  "toonji.com/m/data"
  "toonji.com/m/common"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
  dataResource := TaskResource{}
  err := json.NewDecoder(r.Body).Decode(&dataResouce)
  if err != nil {
    common.WriteError(w,err,"something went wrong",400)
    return
  }
  task = dataResource.Data
  ctx := NewContext()
  defer ctx.Close()
  c := ctx.DBCollection("tasks")
  repo := &data.TaskRepository{c}
  err  = repo.Create(&task)
  if j, err := json.Marshal(TaskResource{Data: task}); err != nil {
    common.WriteError(w,err,"something went wrong",500)
    return
  }
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type","application/json")
  w.Write(j)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
  ctx := NewContext()
  defer ctx.Close()
  c := ctx.DBCollection("tasks")
  repo := data.TaskRepository{&c}
  tasks := repo.GetAll()
  j, err := json.Marshal(TasksResource{Data: tasks})
  if err != nil {
    common.DisplayAppError(w,err,"something went wrong",500)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Write(j)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    v := mux.Var(r)
    id := v["id"]
    ctx := NewContext()
    defer ctx.Close()
    c := ctx.DBCollection("tasks")
    repo := data.TaskRepository{&c}
    task,err := repo.GetById(id)
    if err != nil {
      if err == mgo.ErrNotFound{
        w.WriteHeader(http.StatusNoteContent)
      }else {
        common.WriteError(w,err,"something went wrong",400)
      }
      return
    }
    j, err := json.Marshall(&task)
    if err != nil {
      w.WriteError(w,err,"something went wrong",500)
      return
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOk)
    w.Write(j)
}

func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
// Get id from the incoming url
    vars := mux.Vars(r)
    user := vars["id"]
    context := NewContext()
    defer context. Close()
    c := context.DbCollection("tasks")
    repo := &data.TaskRepository{c}
    tasks := repo.GetByUser(user)
    j, err := json.Marshal(TasksResource{Data: tasks})
    if err != nil {
        common.DisplayAppError(w,err,"An unexpected error has occurred",500)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    // Get id from the incoming url
    vars := mux.Vars(r)
    id := bson.ObjectIdHex(vars["id"])
    var dataResource TaskResource
    // Decode the incoming Task json
    err := json.NewDecoder(r.Body).Decode(&dataResource)
    if err != nil {
        common.DisplayAppError(w,err,"Something went wrong",500)
        return
    }
    task := &dataResource.Data
    task.Id = id
    context := NewContext()
    defer context.Close()
    c := context.DbCollection("tasks")
    repo := &data.TaskRepository{c}
    if err := repo.Update(task); err != nil {
        common.DisplayAppError(w,err,"An unexpected error has occurred",500)
        return
     }else{
        w.WriteHeader(http.StatusNoContent)
    }
}


func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    context := NewContext()
    defer context.Close()
    c := context.DbCollection("tasks")
    repo := &data.TaskRepository{c}
    // Delete an existing Task document
    err := repo.Delete(id)
    if err != nil {
        common.DisplayAppError(w,err,"An unexpected error has occurred",500)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
