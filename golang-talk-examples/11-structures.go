package main

import "fmt"
import "errors"

// Define exported struct
type User struct {
    ID int // public field
    name, password string // private field
}

// Embed User into Admin struct
type Admin struct {
    User
    Privileges []string
}

// Function that operates on object of type User
func (u User) Login(user, password string) (error, bool) {
    if user == u.name && password == u.password {
        return nil, true
    } else {
        return errors.New("Wrong name or password"), false
    }
}

func main () {
    // Create object of type User
    ziomek := User{1, "Ziomek", "123qwe"}

    // Change object field
    ziomek.ID = 2
    err, _ := ziomek.Login("Ziomek", "123qwe")
    if err != nil {
        fmt.Println("zdichu", err.Error())
    }
}