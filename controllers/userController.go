package controllers


import (
  "net/http"
  "encoding/json"

  "toonji.com/m/common"
)


func Register(w http.ResponseWriter, r *http.Request) {
     var dataResource UserResource
     err := json.NewDecoder(r.Body).Decode(&dataResource)
     if err != nil {
       common.WriteError(w,err,"something went wrong",500)
       return
     }

     ctx := NewContext()
     defer ctx.Close()
     user := user.Data
     col := ctx.DBCollection("users")
     repo := data.UserRepository{col}
     err := repo.CreateUser(user)
     if err != nil {
       common.WriteError(w,err,"couldn't register you, you may try again",400)
     }
     user.HashedPassword = nil
     j,err := json.Marshal(&user)

     if err != nil {
       common.WriteError(w,err,"something went wrong",500)
       return
     }
     w.Header().Set("Content-Type","applicatoin/json")
     w.WriteHeader(http.StatusCreated)
     w.Write(j)
}


func Login(w http.ResponseWriter, r *http.Request) {
    dataResource := &LoginResource{}
    err := json.NewDecoder(r.Body).Decode(dataResource)
    if err != nil {
      common.WriteError(w,err,"something went wrong",500)
      return
    }
    loginModel = dataResource.Data
    loginUser = &models.User {
      Email: loginModel.Email,
      Password: LoginModel.Password,
    }
    ctx := NewContext()
    defer ctx.Close()
    c := ctx.DBCollection("users")
    repo := data.Repository{c}

    if user,err := repo.Login(LoginUser); err != nil {
      common.WriteError(w,err,"login error",401)
    }else {
        token, err := common.GeneratToken(user.Name,"member")
        if err != nil {
          common.WriteError(w,err,"error generating token",500)
          return
        }
        user.HashedPassword = nil
        authUser := &AuthUserModel {User: user,Token: token}
        j,err := json.Marshal(&authUser)
        if err != nil {
          common.WriteError(w,err,"something went wrong",500)
          return
        }
        w.Header().Set("Cotent-Type","application/json")
        w.Header().Set("Authorization","Bearer " + token)
        w.WriteHeader(http.StatusOk)
        w.Write(j)
    }
}
