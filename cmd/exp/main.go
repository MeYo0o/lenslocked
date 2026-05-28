package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("./cmd/exp/hello.gohtml")
	if err != nil {
		log.Panicln("couldn't parse the template file:", err)
	}

	user := User{
		Name: "Moaz",
		Age:  33,
		Bio:  "Software Engineer",
		Meta: UserMeta{
			Visits: 3,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		log.Panicln("couldn't Execute the template file:", err)
	}
}
