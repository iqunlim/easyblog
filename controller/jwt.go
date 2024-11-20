package controller

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/iqunlim/easyblog/config"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"github.com/iqunlim/easyblog/service"
)

var (
	identityKey = "id"
	roleKey = "Role"
)

type Identity struct {
	Username string
	Role string
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
		Realm:       config.JWTRealm,
		IdentityKey: identityKey,
		// Login Workflow
		Authenticator: AuthenticateUser(),
		PayloadFunc: LoginPayLoad(),
		LoginResponse: LoginResponse(),

		// Middleware Workflow
		IdentityHandler: identityHandler(),
		Authorizator:    Authorize(),

		//Logout Workflow

		LogoutResponse: LogoutResponse(),

		//Refresh Workflow
		//RefreshResponse: RefreshResponse(),

		// Extra
		Unauthorized:    Unauthorized(),
		// Tokens
		Key:         []byte(config.JWTKey), //TODO: pem key
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 12,
		TimeFunc:      time.Now,
		TokenLookup: "header:Authorization,cookie:jwt",
		TokenHeadName: "Bearer",
		// Cookies
		SendCookie: true,
		SecureCookie: false, // Set to true for http in prod
		CookieHTTPOnly: true, // Remove in prod, we do not want client side to access cookies in prod
		CookieDomain: config.Domain, //Set to your domain in prod
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

		db := model.GetDB()
		
		// TODO: This IS TERRIBLE
		u := repository.NewUserRepository(db)
		us := service.NewUserService(u)
		userInfo, err := us.Verify(newUser)
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
				roleKey: v.Role,
			}
		}
		return jwt.MapClaims{}
	}
} 


func LoginResponse() func(c *gin.Context, code int, message string, time time.Time) {
	return func(c *gin.Context, code int, message string, time time.Time) {
		c.Redirect(302, "/admin/")
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &Identity{
			Username: claims[identityKey].(string),
			Role: claims[roleKey].(string),
		}
	}
}


func Authorize() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		// Here is where we make the magic happen
		// Probably have some sort of token in identity or userinfo
		if v, ok := data.(*Identity); ok && v.Role == "admin" {
			return true
		} 
		return false
	}
}

func LogoutResponse() func(c *gin.Context, code int) {
	return func(c *gin.Context, code int) {
		c.Redirect(302, "/posts")
	}
}


func Unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string){
		c.Redirect(302, "/login")
	}
}