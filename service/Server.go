package service

import (
	"github.com/go-martini/martini"
)

func NewServer(port string) {
	m := martini.Classic()
	// To get the input name using martini.Params
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})
	// To change the port
	m.RunOnAddr(":"+port)	
}