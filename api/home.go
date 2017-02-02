package api

import "github.com/kataras/iris"

type HomeApi struct {
	*iris.Context
}

func (this HomeApi) Homepage(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{
		"title":   "Gitalent - Show your talent!",
		"message": "Initial",
	})
}
