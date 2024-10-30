package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqunlim/loginexample/model"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {

	var registerInput login 
	if err := c.ShouldBind(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
		return
	}


	u := model.User{
		Username: registerInput.Username,
		Password: registerInput.Password,
	}

	err := u.SaveToDB()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
		return
	}

	c.String(200, "Working")
}



