package controllers

import (
  "gopkg.in/mgo.v2"
  "toonji.com/m/common"
)


type Context struct {
  MongoSession *mgo.Session
}

func(c *Context) Close() {
   c.MongoSession.Close()
}

func (c *Context) DBCollection(name string) *mgo.Collection {
   session := c.MongoSession
   return session.DB(common.AppConfig.Datase).C(name)
}

func NewContext() *Context {
  ctx :=  &Context{
    MongoSession: common.GetSession().Copy(),
  }
  return ctx
}
