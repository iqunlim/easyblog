package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iqunlim/loginexample/model"
)

func handleBlogPostGet(c *gin.Context) {
	// Takes URI of post ID /blog/{id}
	// Takes query parameters ?term=tech
	// This will search the title, category, 
	// and then the content of the markdown file if required
	// Returns BlogReturn
	// Return codes:
	// 200 OK if it is created
	// 404 Not found with err if not
}

func handleBlogPostPost(c *gin.Context) {

	var blog model.BlogPost
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.String(400, "Invalid")
		return
	}
	c.String(200, "Valid")
	// Takes Blog JSON
	// Returns BlogReturn json
	// Return codes:
	// 201 Created if it is created
	// 400 Bad Request with err if not
}

func handleBlogPostUpdate(c *gin.Context) {
	// Takes Blog struct 
	// Returns BlogReturn
	// Return codes:
	// 200 OK if updated
	// 400 Bad request if invalid
	// 404 Not Found if blog not found
}

func handleBlogPostDelete(c *gin.Context) {
	// Takes ID a-la /blog/{id}
	// Returns BlogReturn
	// Return codes:
	// 200 OK if deleted
	// 404 Not Found if blog post not found
}


func handleBlogPostGetAll(c *gin.Context) {
	// Takes NOTHING
	// Returns []BlogReturn
	// Return codes:
	// 200 OK if array 
	// TODO: Consider pagination, etc in the future
}