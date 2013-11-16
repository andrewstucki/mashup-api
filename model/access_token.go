package model

import (
  "github.com/garyburd/redigo/redis"
  "github.com/mashup-cms/mashup-api/globals"
  "github.com/mashup-cms/mashup-api/helpers"
  "log"
  "fmt"
  "strconv"
)

type AccessToken struct {
  UserId int `json:"-"`
  AppId int `json:"-"`
  Scopes []string `json:"-"`
  Token string `json:"token"`
}

func (token *AccessToken) User() (*User) {
  user, err := globals.PostgresConnection.Get(User{}, token.UserId)
  if err != nil {
    log.Fatal("Retrieval Error: ", err)
  }
  return user.(*User)
}

func FindAccessToken(token string) (*AccessToken) {
  key := fmt.Sprintf("t:%s", token)

  connection := globals.RedisPool.Get()
  session, err := redis.Strings(connection.Do("LRANGE", key, 0, -1))
  if err != nil || len(session) < 4 {
    defer connection.Close()  
    return nil  
  }
  scopes := session[2:]
  userId, err := strconv.ParseInt(session[0],0,0)
  if err != nil {
    defer connection.Close()  
    return nil  
  }
  appId, err := strconv.ParseInt(session[1],0,0)
  if err != nil {
    defer connection.Close()  
    return nil
  }
  defer connection.Close()

  return &AccessToken{
    UserId: int(userId),
    AppId: int(appId),
    Scopes: scopes,
    Token: token,
  }
}

func GenerateAccessToken(uId int, appId int) (*AccessToken) {
  //do push to Redis

  scopes := []string{"public", "private"}
  token := helpers.SecureToken()
  key := fmt.Sprintf("t:%s", token)
  reuseKey := fmt.Sprintf("r:%d:%d", uId, appId)
    
  connection := globals.RedisPool.Get()
  connection.Send("MULTI")
  
  connection.Send("RPUSH", key, uId)
  connection.Send("RPUSH", key, appId)
  connection.Send("RPUSH", key, scopes[0])
  connection.Send("RPUSH", key, scopes[1])  
  connection.Send("SET", reuseKey, token)
  
  _, err := connection.Do("EXEC")
  defer connection.Close()
  
  if err != nil {
    log.Printf("Error: %s", err.Error())
    return nil
  }
  
  log.Printf("%s", uId)
  
  return &AccessToken {
    UserId: uId,
    AppId: appId,
    Scopes: scopes,
    Token: token,
  }

  //Handle Expiration
}

func GetReusableAccessToken(uId int, appId int) (*AccessToken) {
  key := fmt.Sprintf("r:%d:%d", uId, appId)
  
  connection := globals.RedisPool.Get()
  token, err := connection.Do("GET", key)
  defer connection.Close()
  
  if err != nil || token == nil {
    return nil
  }
  
  return &AccessToken{
    UserId: uId,
    AppId: appId,
    Scopes: []string{"public","private"},
    Token: string(token.([]byte)),
  }
}