package common


import (
  "log"
  "encoding/json"
  "net/http"
)


type (
  appError struct {
     Error      string `json:"error"`
     Message    string `json:"message"`
     HTTPStatus int `json:"status"`
  }

  errorResource struct {
     Data appError `json:"data"`
  }
)

func WriteError(w http.ResponseWriter, err error, message string, status int) {
    errRouse := errorResource {
         appError {
           Error: err.Error(),
           Message: message,
           HTTPStatus: status,
         },
    }
    w.Header().Set("content-type","applicatoin/json; char=utf-8")
    w.WriteHeader(status)
    body, err := json.Marshal(&errRouse)
    if err != nil {
      log.Fatalln("something went wrong")
    }
    w.Write(body)
}
