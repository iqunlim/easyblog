package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqunlim/easyblog/config"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"github.com/iqunlim/easyblog/service"
)


type WebHandler interface {
	LoginWebHandler(*gin.Context)
	RegisterWebHandler(*gin.Context)
	PostsWebHandler(*gin.Context)
	PostsByIDWebHandler(*gin.Context)
	AdminDashboardHandler(*gin.Context)
	AdminEditPostHandler(*gin.Context)
}


type WebHandlerImpl struct {

	blogservice service.BlogService
}

func NewWebHandler(blogservice service.BlogService) WebHandler {
	return &WebHandlerImpl{
		blogservice: blogservice,
	}
}


func (w *WebHandlerImpl) LoginWebHandler(c *gin.Context) {

	// Some stuff to check if user has already logged in
	// TODO: Better method than checking the cookie?
	// After some testing, this may be the best way
	// Claims are not sent from the jwt middleware to this
	// Endpoint.
	// If the cookie is expired, this will return an error as well
	if !isLoggedIn(c) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	} else {
		// Oops already logged in, why are you re-accessing?
		c.Redirect(302, "/posts")
	}
}

func (w *WebHandlerImpl) RegisterWebHandler(c *gin.Context) {
	if config.RegisterAllowed == "true" {
		c.HTML(200, "register.html", gin.H{"title": "Register Admin Account"})
	} else {
		c.Redirect(302, "/posts")
	}
	//c.HTML(500, "error.html", gin.H{"err": "Not Implemeneted"})
}

func (w *WebHandlerImpl) PostsWebHandler(c *gin.Context) {

	//term := c.DefaultQuery("term", "NONE")
	data := gin.H{}
	if isLoggedIn(c) {
		data["auth"] = "true"
	}

	posts, err := w.blogservice.GetAllNoContent(c)
	if err != nil {
		c.HTML(500, "error.html", gin.H{"err": "Internal Server Error"})
	}
	data["title"] = "All Posts"
	data["posts"] = posts
	c.HTML(http.StatusOK, "index.html", data)

	// TODO: 
	// Stage 2: Consider pagination
}

func (w *WebHandlerImpl) PostsByIDWebHandler(c *gin.Context) {

	var id blogid
	var data = gin.H{}
	if isLoggedIn(c) {
		data["auth"] = "true"
	}
	if err := c.ShouldBindUri(&id); err != nil {
		c.HTML(400, "error.html", gin.H{"err": "Malformed Request"})
		return
	}

	ret, err := w.blogservice.GetByID(c, id.ID, true)
	if err != nil {
		var be *repository.NotFoundError
		if errors.As(err, &be) {
			c.Redirect(302, "/posts")
			return
		} else {
			c.HTML(404, "error.html", gin.H{"err": "Post Not Found"})
			return
		}
	}

	data["post"] = ret;
	data["title"] = ret.Title


	c.HTML(http.StatusOK, "post.html",data)
}

func (w *WebHandlerImpl) AdminDashboardHandler(c *gin.Context) {

	ret, err := w.blogservice.GetAllNoContent(c)
	if err != nil {
		var be *repository.NotFoundError
		if errors.As(err, &be) {
			c.HTML(200, "admin.html", gin.H{"posts": nil, "auth": "true"})
			return
		} else {
			c.HTML(500, "error.html", gin.H{"err": "Internal Server Error"})
			return
		}
	}

	c.HTML(200, "admin.html", gin.H{"posts": ret, "auth":"true"})
}

func (w *WebHandlerImpl) AdminEditPostHandler(c *gin.Context) {

	var id blogid
	if err := c.ShouldBindUri(&id); err != nil {
		c.HTML(200, "edit.html", gin.H{"title": "New Post", "post": model.BlogPost{}, "auth": "true"})
		return
	}

	post, err := w.blogservice.GetByID(c, id.ID, false)

	if err != nil {
		var be *repository.NotFoundError
		if errors.As(err, &be) {
			c.HTML(200, "edit.html", gin.H{"title": "New Post", "post": model.BlogPost{}, "auth":"true"})
			return
		} else {
			c.HTML(500, "error.html", gin.H{"err": "Internal Server Error"})
			return
		}
	}
	c.HTML(200, "edit.html", gin.H{"title": "Edit Post", "post": post, "auth":"true"})
}