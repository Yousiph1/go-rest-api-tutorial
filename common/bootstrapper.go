package common

import (
  "log"
  "os"
  "encoding/json"
)

type configuration struct {
  Server, MongoDBHost, DBUser,DBPwd,Database string
}

func initConfig() {
  loadConfig()
}

var AppConfig *configuration

func loadConfig() {
   file,err := os.Open("common/config.json")
   defer file.Close()
   if err != nil {
     log.Fatalf("ERROR: %s\n",err)
   }
   decoder := json.NewDecoder(file)
   AppConfig = &configuration{}
   err = decoder.Decode(AppConfig)
   if err != nil {
      log.Fatalf("ERROR: %s\n",err)
   }
}

func StartUp() {
   initConfig()
   initKeys()
   createDbSession()
   addIndexes()
}
