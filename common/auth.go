package common

import (
  "io/ioutil"
  "log"
  "time"
  "net/http"

  "github.com/dgrijalva/jwt-go"
)

const (
  publicKey = "/keys/app.rsa.pub"
  privateKey = "/keys/app.rsa"
)

var (
  verifyKey, signKey []byte
)

func initKeys() {
  var err error
  verifyKey,err = ioutil.ReadFile(publicKey)
  if err != nil {
    log.Fatalf("[initKeys: read public key file error] %s\n",err)
    panic(err)
  }
  signKey,err = ioutil.ReadFile(privateKey)
  if err != nil {
    log.Fatalf("]initKeys: read private key file error] %s\n",err)
  }
}

func GenerateToken(name, role string) (string,error) {
   t := jwt.New(jwt.GetSigningMethod("RS256"))

   t.Claims["iss"] = "admin"
   t.Claims["UserInfo"] = struct {
     Name, Role string
   }{name,role}
   t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
   tokenString, err := t.SignedString(signKey)
   if err != nil {
     return "",err
   }
   return tokenString, nil
}
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    token, err := jwt.Parse(r,func (token *jwt.Token)(interface{},error){
      return verifyKey, nil
    })

    if err != nil {
      switch err.(type) {
      case jwt.ValidationError:
          vErr := err.(*jwt.ValidationError)
          switch vErr.Errors {
          case jwt.ValidationErrorExpired:
            WriteError(
              w,
              err,
              "token has expired, get a new token",
              401,
            )
            return
          default:
          WriteError(
            w,
            err,
            "error while parsing token",
            500,
          )
          return
          }
      default :
          WriteError(
            w,
            err,
            "error while parsing access token",
            500,
          )
          return
      }
    }
    if token.Valid {
      next(w,r)
    }else {
      WriteError(
        w,
        err,
        "Invalid access token",
        401,
      )
    }
}
