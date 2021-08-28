package controllers


import (
  "toonji.com/m/models"
)

type (
  UserResource struct {
    Data  model.User `json:"data"`
  }

  LoginResource struct {
    Data  LoginModel `json:"data"`
  }

  LoginModel struct {
    Name     string `json:"name"`
    Password string `json:"password"`
  }

  AuthUserResource struct {
    Data AuthUserModel `json:"data"`
  }

  AuthUserModel struct {
    User  model.User `json:"user"`
    Token string     `json:"token"`
  }
)


type (
  TaskResource struct {
    Data  model.Task  `json:"data"`
  }

  TasksResource struct {
    Data  []model.Task `json:"data"`
  }
)

type (

    NoteResource struct {
      Data NoteModel `json:"data"`
    }

    NotesResource struct {
      Data []models.TaskNote `json:"data"`
    }
    
    NoteModel struct {
      TaskId string `json:"taskid"`
      Description string `json:"description"`
    }
)
