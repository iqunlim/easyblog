package controller

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/iqunlim/easyblog/config"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"github.com/iqunlim/easyblog/service"
)

//Standard App response structs and functions
// Example usage
//app := Gin{C: c}
//app.Response(200, true, "Success", nil)
type APIResponse struct {
	Code		int    `json:"code"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	MessageBody any    `json:"data,omitempty"`
}

// Used to create an app object
type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, success bool, messagetext string, data any) {
	g.C.JSON(httpCode, APIResponse{
		Code: 		 httpCode,
		Success:     success,
		Message:     messagetext,
		MessageBody: data,
	})
	return
}

// Cache Control for static files
func staticCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Apply the Cache-Control header to the static files
		if strings.HasPrefix(c.Request.URL.Path, "/static/") {
			c.Header("Cache-Control", "public, max-age=86400")
		}
		// Continue to the next middleware or handler
		c.Next()
	}
}
// Create the engine
func CreateAPI() (*gin.Engine, error) {

	r := gin.Default()

	//TODO: Fix this when prod
	r.SetTrustedProxies(nil)

	authMiddleware, err := jwt.New(initParams())
	if err != nil {
		return nil, err
	}
	// Cache Middleware
	r.Use(staticCacheMiddleware())

	//CORS middleware
	//r.Use(cors.New(config.CORSConfig))

	// Implement the jwt middleware
	r.Use(handlerMiddleWare(authMiddleware))

	// Set up Handlers, Services, and Repositories
	db := model.GetDB()
	userRep := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRep)
	authHandler := NewAuthHandler(userSvc)

	blogRep := repository.NewBlogRepository(db)
	blogSvc := service.NewBlogService(blogRep)

	imageRep := repository.NewImageRepositoryLocalhost(config.ProjectRoot + "/static/files/")
	imageSvc := service.NewImageService(imageRep)


	blogHandler := NewBlogHandler(blogSvc, imageSvc)

	webHandler := NewWebHandler(blogSvc)

	
	// Template functions for mutating variables in the template
	r.SetFuncMap(template.FuncMap{
		"formatAsHTML": formatAsHTML,
		"formatTimeToString": formatTimeToString,
	})

	// Static file handling for the webserver portion
	r.LoadHTMLGlob("static/template/*.html")
	r.Static("/static", "./static")
	r.MaxMultipartMemory = 20 << 20 //20 MB


	// API Endpoints
	api := r.Group("/api/v1")
	api.POST("/register", authHandler.RegisterHandler)
	api.POST("/login", authMiddleware.LoginHandler)
	api.GET("/refresh", authMiddleware.RefreshHandler)
	api.GET("/posts", blogHandler.handleBlogPostGetAll)
	api.GET("/posts/:id", blogHandler.handleBlogPostGet)
	// HTML endpoints
	r.GET("/", webHandler.PostsWebHandler)
	r.GET("/home", webHandler.PostsWebHandler)
	r.GET("/posts", webHandler.PostsWebHandler)
	r.GET("/posts/:id", webHandler.PostsByIDWebHandler)
	r.GET("/login", webHandler.LoginWebHandler)
	r.GET("/logout", authMiddleware.LogoutHandler)
	r.GET("/register", webHandler.RegisterWebHandler)
	// Admin endpoints
	requiresAuth := r.Group("/admin", authMiddleware.MiddlewareFunc())
	requiresAuth.GET("/", webHandler.AdminDashboardHandler)
	requiresAuth.GET("/edit", webHandler.AdminEditPostHandler)
	requiresAuth.GET("/edit/:id", webHandler.AdminEditPostHandler)

	// Admin API endpoints
	requiresAuth.DELETE("/posts/:id", blogHandler.handleBlogPostDelete)
	requiresAuth.POST("/posts", blogHandler.handleBlogPostPost)
	requiresAuth.PUT("/posts/:id", blogHandler.handleBlogPostUpdate)
	r.POST("/posts/image", blogHandler.handleBlogImageUpload)
	return r, nil
}

func formatAsHTML(s string) template.HTML {
	return template.HTML(s)
}


func formatTimeToString(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%d", month, day, year)
}


func isLoggedIn(c *gin.Context) bool {
	_, e := c.Cookie("jwt")
	if e != nil {
		return false 
	}
	return true
}