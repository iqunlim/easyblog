package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"github.com/iqunlim/easyblog/service"
)

type BlogHandler interface {
	handleBlogPostGet(*gin.Context)
	handleBlogPostGetAll(*gin.Context)
	handleBlogPostPost(*gin.Context)
	handleBlogPostUpdate(*gin.Context)
	handleBlogPostDelete(*gin.Context)
	handleBlogImageUpload(*gin.Context)
}

type BlogHandlerImpl struct {
	blogservice service.BlogService
	imageservice service.ImageHandlerService
}

func NewBlogHandler(blogservice service.BlogService, imageservice service.ImageHandlerService) BlogHandler {
	return &BlogHandlerImpl{
		blogservice: blogservice,
		imageservice: imageservice,
	}
}

// Struct for handling the URI
type blogid struct {
	ID int `json:"id" uri:"id" binding:"required"`
} 


func (b *BlogHandlerImpl) handleBlogPostGet(c *gin.Context) {


	app := Gin{C: c}
	var id blogid
	if err := c.ShouldBindUri(&id); err != nil {
		app.Response(400, false, "Malformed Request", nil)
		return
	}

	ret, err := b.blogservice.GetByID(c, id.ID, true)
	if err != nil {
		var be *repository.NotFoundError
		if errors.As(err, &be) {
			app.Response(404, false, "Post not Found", nil)
			return
		} else {
			app.Response(500, false, "Internal Server Error", nil)
			//TODO: Log Error
			return
		}
	}
	app.Response(200, true, "Success", ret)
}

func (b *BlogHandlerImpl) handleBlogPostPost(c *gin.Context) {

	app := Gin{C: c}
	var blog model.BlogPost

	if err := c.ShouldBind(&blog); err != nil {
		app.Response(400, false, "Malformed Request", nil)
		return
	}

	res, err := b.blogservice.Post(c, &blog)
	if err != nil {
		app.Response(500, false, "Internal Server Error", nil)
		return
	}

	app.Response(201, true, "Success", res)
}

func (b *BlogHandlerImpl) handleBlogPostUpdate(c *gin.Context) {

	//TODO: Only return the parts of the blogpost that are updated and nothing else

	app := Gin{C: c}
	var blog model.BlogPost
	var id blogid
	if err := c.ShouldBindUri(&id); err != nil {
		app.Response(400, false, "Malformed Request", nil)
		return
	}
	if err := c.ShouldBind(&blog); err != nil {
		app.Response(400, false, "Malformed Request", nil)
		return
	}
	if err := b.blogservice.Update(c, id.ID, &blog); err != nil {
		var be *repository.NotFoundError
		if errors.As(err, &be) {
			app.Response(404, false, fmt.Sprintf("Post %d not Found", id.ID), nil)
			return
		} else {
			app.Response(500, false, "Internal Server Error", nil)
			return
		}
	}
	app.Response(200, true, "Success", blog)
}

func (b *BlogHandlerImpl) handleBlogPostDelete(c *gin.Context) {

	app := Gin{C: c}
	var id blogid
	if err := c.ShouldBindUri(&id); err != nil {
		app.Response(400, false, "Malformed Request", nil)
		return
	}
	if err := b.blogservice.Delete(c, id.ID); err != nil {
		var be *repository.NotFoundError
		if errors.As(err, &be) {
			app.Response(404, false, fmt.Sprintf("Post %d not Found", id.ID), nil)
			return
		} else {
			app.Response(500, false, "Internal Server Error", nil)
			return
		}
	}
	app.Response(200, true, fmt.Sprintf("Deleted post %d successfully", id.ID), nil)
}


func (b *BlogHandlerImpl) handleBlogPostGetAll(c *gin.Context) {


	term := c.DefaultQuery("term", "NONE")

	app := Gin{C: c}
	posts, err := b.blogservice.GetAll(c, term, true)
	if err != nil {
		app.Response(500, false, "Internal Server Error", nil)
	}
	app.Response(200, true, "Success", posts)
	// TODO: 
	// Stage 2: Consider pagination
}

func (b *BlogHandlerImpl) handleBlogImageUpload(c *gin.Context) {


	app := Gin{C: c}
	file, err := c.FormFile("file")
	if err != nil {
		app.Response(400, false, "Malformed Request", nil)
	}
	fileBody, err := file.Open()
	defer fileBody.Close()
	if err != nil {
		app.Response(400, false, "Malformed Request", nil)
	}

	ret, err := b.imageservice.Upload(c, fileBody, file)
	if err != nil {
		app.Response(500, false, "Internal Server Error", err.Error())
	}

	app.Response(200, true, "Success", gin.H{ "filepath": ret})

}