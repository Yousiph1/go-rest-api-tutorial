package models

import (
  "time"
  "gopkg.in/mgo.v2/bson"
)

type (

  // NOTE: do not store passwords in plain text
  User struct {
    ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
    FirstName      string        `json:"firstname"`
    LastName       string        `json:"lastname"`
    Email          string        `json:"email"`
    Password       string        `json:"password"`
    HashedPassword []byte        `json:"hashedpassword,omitempty"`

  }

  Task struct {
    ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
    CreatedBy     string        `json:"createdby"`
    Name          string        `json:"name"`
    Description   string        `json:"description"`
    CreatedOn     time.Time     `json:"createdon"`
    Due           time.Time     `json:"due"`
    Status        string        `json:"status"`
    Tags          []string      `json:"tags"`
  }

  TaskNote struct {
   ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
   TaskID        bson.ObjectId `josn:"taskid"`
   Description   string        `json:"description"`
   CreatedOn     time.Time     `json:"createdon"`
  }
)
