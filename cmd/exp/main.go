package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Panicln("couldn't parse the template file")
	}

	user := User{
		Name: "Moaz",
		Age:  33,
		Meta: UserMeta{
			Visits: 3,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		log.Panicln("couldn't Execute the template file")
	}
}
