package controllers

import (
	"github.com/astaxie/beego"
)

type Server struct {
	beego.Controller
}

func (server *Server) Initialize() {
	beego.Router("/sort/:list", server)
}

func (server *Server) Run() {
	beego.Run()
}