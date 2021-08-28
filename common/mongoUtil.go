package common

import (
  "log"
  "time"
  "gopkg.in/mgo.v2"
)

var session *mgo.Session


func GetSession() *mgo.Session {
   var err error
   session, err = mgo.DialWithInfo(&mgo.DialInfo{
     Addrs:    []string{AppConfig.MongoDBHost},
     Username: AppConfig.DBUser,
     Password: AppConfig.DBPwd,
     Timeout:  60 * time.Second,
   })
   if err != nil {
     log.Fatalf("[GetSession]: %s\n",err)
   }
   return session
}


func createDbSession() {
       var err error
      session, err = mgo.DialWithInfo(&mgo.DialInfo{
          Addrs: []string{AppConfig.MongoDBHost},
          Username: AppConfig.DBUser,
          Password: AppConfig.DBPwd,
          Timeout: 60 * time.Second,
      })
      if err != nil {
          log. Fatalf("[createDbSession]: %s\n", err)
      }
}

func failOnError(err error,msg string) {
  if err != nil {
    log.Fatalf("%s %s\n",msg,err)
  }
}

func addIndexes() {
   userIndex := mgo.Index{
     Key:        []string{"email"},
     Unique:     true,
     Background: true,
     Sparse:     true,
   }
   taskIndex := mgo.Index{
     Key:        []string{"createdby"},
     Unique:     false,
     Background: true,
     Sparse:     true,
   }
   noteIndex := mgo.Index{
      Key:        []string{"taskid"},
      Unique:     false,
      Background: true,
      Sparse:     true,
   }

   session := GetSession().Copy()
   defer session.Close()
   userCol := session.DB(AppConfig.Database).C("users")
   taskCol := session.DB(AppConfig.Database).C("tasks")
   noteCol := session.DB(AppConfig.Database).C("notes")

   err := userCol.EnsureIndex(userIndex)
   failOnError(err,"[user addIndex]:")
   err = taskCol.EnsureIndex(taskIndex)
   failOnError(err,"[task addIndex]:")
   err = noteCol.EnsureIndex(noteIndex)
   failOnError(err,"[note addIndxe]:")
}
