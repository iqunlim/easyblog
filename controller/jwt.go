package controller

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/iqunlim/loginexample/model"
)

var identityKey = "id"

type Identity struct {
	Username string
	AuthKey string
}
  
func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
	  	err := authMiddleware.MiddlewareInit()
	  	if err != nil {
			log.Fatal("authMiddleware.MiddlewareInit:" + err.Error())
	 	}
	}
}

func initParams() *jwt.GinJWTMiddleware {

	// Provided Entrypoints:
	// LoginHandler: Starts Login Workflow
	// MiddlewareFunc() Starts middleware Workflow
	// LogoutHandler: Starts logout workflow
	// RefreshHandler: Starts Refresh workflow
  
	return &jwt.GinJWTMiddleware{
		Realm:       "IQBlog",
		IdentityKey: identityKey,
		// Login Workflow
		Authenticator: AuthenticateUser(),
		PayloadFunc: LoginPayLoad(),
		//LoginResponse: ,

		// Middleware Workflow
		IdentityHandler: identityHandler(),
		Authorizator:    Authorize(),

		//Logout Workflow
		//LogoutResponse: ,

		//Refresh Workflow
		// RefreshResponse: ,

		// Extra
		Unauthorized:    Unauthorized(),
		// Tokens
		Key:         []byte("SDFJKHSDHFJKSHKJDFHKJSDF"), //TODO: pem key
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		TimeFunc:      time.Now,
		TokenLookup: "header:Authorization,cookie:jwt",
		TokenHeadName: "Bearer",
		// Cookies
		SendCookie: true,
		SecureCookie: false, // Set to true for http in prod
		CookieHTTPOnly: true, // Remove in prod
		CookieDomain: "localhost", //Set to your domain in prod
	}
}

func AuthenticateUser() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login 
		if err := c.ShouldBind(&loginVals); err != nil {
			return nil, jwt.ErrMissingLoginValues
		}

		newUser := &model.User{
			Username: loginVals.Username,
			Password: loginVals.Password,
		}

		userInfo, err := model.VerifyLogin(newUser)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		return userInfo, nil		
	}
}

func LoginPayLoad() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*model.User); ok {
			return jwt.MapClaims{
				identityKey: v.Username,

				
			}
		}
		return jwt.MapClaims{}
	}
} 

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &Identity{
			Username: claims[identityKey].(string),
		}
	}
}


func Authorize() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		// Here is where we make the magic happen
		// Probably haave some sort of token in identity or userinfo
		if v, ok := data.(*Identity); ok && v.Username == "admin" {
			return true
		} 
		return false
	}
}


func Unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string){
		c.JSON(code, gin.H{
			"code": code,
			"message": message,
		})
	}
}

// Test handler
func helloHandler(c *gin.Context) {
  claims := jwt.ExtractClaims(c)
  user, _ := c.Get(identityKey)
  fmt.Println(user)
  log.Println(claims)
}