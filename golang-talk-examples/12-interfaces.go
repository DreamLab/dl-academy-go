package main

import "fmt"
import "errors"

// define method for login
type User interface {
  Login(user, password string) (error, bool)
}

type Admin struct {
   Privileges []string
   Name, Pass string
}

// function that operates on object of type Admin
func (a Admin) Login(user,passwd string) (error, bool) {
    if user == a.Name && passwd == a.Pass {
       return nil, true
    } else {
       return errors.New("Wrong login or password"), false
    }
}

type SuperAdmin struct {
   Name, Pass string
}

// function that operates on object of type User
func (a SuperAdmin) Login(user, passwd string) (error, bool) {
    if user == a.Name && passwd == a.Pass {
       return nil, true
    } else {
       return errors.New("Wrong login or password"), false
    }
}

// to function login we can pass Admin and SuperAdmin
func LoginUser(u User) {
    err, _ := u.Login("admin", "admin")

    if err != nil {
        fmt.Println(err.Error())
    }
}

func main () {
    user := Admin{Name: "admin", Pass: "admin"}
    superUser := SuperAdmin{"root", "root"}

    LoginUser(user)
    LoginUser(superUser)
}
