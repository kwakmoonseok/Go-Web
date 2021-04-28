package main

import (
	"html/template"
	"os"
)


type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name:"kinder", Email:"example@gmail.com", Age:23}
	user2 := User{Name:"aaa", Email:"aaa@gmail.com", Age:40}
	users := []User{user, user2}
	tmpl, err := template.New("Tmpl1").ParseFiles("templetes/tmpl1.tmpl", "templetes/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}