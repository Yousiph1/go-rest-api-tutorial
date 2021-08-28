package main

import (
  "net/http"
  "log"

  "github.com/codegangsta/negroni"
  "toonji.com/m/common"
  "toonji.com/m/routers"
)


func main() {
  common.StartUp()
  router := routers.InitRoutes()
  n := negroni.Classic()
  n.UserHandler(router)
  server := &http.Server{
    Addr: common.AppCofig.Server,
    Handler: n,
  }
  server.ListenAndServe()
  log.Println("server started")
}
