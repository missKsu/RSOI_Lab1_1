package main

import "github.com/go-martini/martini"

func main(){
	m := martini.Classic()
	m.Get("/app",func() string{
		return "Hello World!"
		})
	m.Run()
}