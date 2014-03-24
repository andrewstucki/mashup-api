package model

import(
  "code.google.com/p/go.crypto/bcrypt"

  "time"
  "github.com/mashup-cms/mashup-api/globals"
)

type UserService struct {
  User *User `json:"user"`
}

type User struct {
  Id int `db:"id" json:"id"`
  Name NullString `db:"name" json:"name,omitempty"`
  Email NullString `db:"email" json:"email,omitempty"`
  Login string `db:"login" json:"login,omitempty"`
  
  EncryptedPassword string `db:"encrypted_password" json:"-"`
  CreatedAt time.Time `db:"created_at" json:"-"`
  UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (user User) GetId() (int){
  return user.Id
}

func (user User) GetName() (string){
  if user.Name.Valid {
    return user.Name.String
  } else {
    return ""
  }
}

func (user User) GetEmail() (string){
  if user.Email.Valid {
    return user.Email.String
  } else {
    return ""
  }
}

func (user *User) PreInsert(interface{}) error {
    user.CreatedAt = time.Now().UTC()
    user.UpdatedAt = user.CreatedAt
    return nil
}

func (user *User) PreUpdate(interface{}) error {
    user.UpdatedAt = time.Now().UTC()
    return nil
}

func (user *User) GetService() interface{} {
  return UserService{User:user}
}

func (user *User) CheckPassword(password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
  if err != nil {
    return false
  }
  return true
}

func CreateUser(username string, password string) (User, error) {
  pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
  user := User{Login:username, EncryptedPassword:string(pass)}
  if err == nil {
    err = globals.PostgresConnection.Insert(&user)
  }
  return user, err
}