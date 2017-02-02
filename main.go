package main

import (
	"gitalent.com/backend/api"
	"github.com/kataras/iris"
)

func main() {
	// iris.StaticWeb("/dist", "../talent")
	// iris.Config.Render.Template.IsDevelopment = true
	registerAPI()

	//iris.Get("/*file", iris.StaticHandler("/", "../talent/dist", false, true), api)

	iris.Listen("localhost:8000")
}

func registerAPI() {
	home := new(api.HomeApi)

	iris.Get("/v1/home", home.Homepage)
}
