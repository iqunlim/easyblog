package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/service"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


type AuthHandler interface {
	RegisterHandler(*gin.Context)
}

type AuthHandlerImpl struct {
	userservice service.UserService
}

func NewAuthHandler(userservice service.UserService) AuthHandler {
	return &AuthHandlerImpl{
		userservice: userservice,
	}
}

func (a *AuthHandlerImpl) RegisterHandler(c *gin.Context) {

	app := Gin{C: c}
	var u model.User 
	if err := c.ShouldBind(&u); err != nil {
		app.Response(400, false, "Malformed Request", nil)
		return
	}
	err := a.userservice.Register(&u)

	if err != nil {
		var se *service.UserExistsError
		if errors.As(err, &se) {
			app.Response(500, false, "Username exists", nil)
			return
		}
		app.Response(500, false, "Internal Server Error", nil)
		return
	}

	c.Redirect(302, "/login")
}
