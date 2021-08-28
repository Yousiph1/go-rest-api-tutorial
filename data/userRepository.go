package data

import (
  "golang.org/x/crypto/bcrypt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "toonji.com/m/models"
)

type UserRepository struct {
   C *mgo.Collection
}

func (r *UserRepository) CreateUser(u models.User) error {
  _id := bson.NewObjectId()
  u.ID = _id
  hp,err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  u.HashedPassword = hp
  u.Password = ""
  r.C.Insert(&u)
  return nil
}

func (r *UserRepository) Login(u models.User) (user models.User, err error) {

   err = r.C.Find(bson.M{"email":u.Email}).One(&user)
   if err != nil {
     return
   }
   err = bcrypt.CompareHashAndPassword(user.HashedPassword,[]byte(u.Password))
   if err != nil {
     user = models.User{}
     return
   }
  return
}
