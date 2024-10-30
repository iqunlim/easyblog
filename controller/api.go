package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/iqunlim/loginexample/model"
)

func CreateAPI() error {

	r := gin.Default()

	//TODO: Fix this when prod
	r.SetTrustedProxies(nil)
	var err error
	model.DBSingleton, err = model.CreateDB()
	if err != nil {
		panic(err)
	}
	authMiddleware, err := jwt.New(initParams())
	if err != nil {
		panic(err)
	}

	r.Use(handlerMiddleWare(authMiddleware))

	r.POST("/register", RegisterHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/logout", authMiddleware.LogoutHandler)
	requiresAuth := r.Group("/admin", authMiddleware.MiddlewareFunc())
	requiresAuth.GET("/hello", helloHandler)
	r.POST("/new", handleBlogPostPost)


	if err := r.Run(":8080"); err != nil {
		return err
	}
	return nil
}