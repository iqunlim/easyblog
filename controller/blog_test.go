package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"github.com/iqunlim/easyblog/service"
	mock "github.com/stretchr/testify/mock"
)

var (
	FakeBlogPost = &model.BlogPost{
		Title:    "Test",
		Content:  "TestTestTest",
		Category: "Test",
		Tags:     model.Tags{"Tech", "Test"},
	}

	FakeBlogPost2 = &model.BlogPost{
		Title:    "Test2",
		Content:  "TestTest",
		Category: "Test2",
		Tags:     model.Tags{"Test"},
	}

)

type APIResponseB struct {
	MessageBody model.BlogPost `json:"data,omitempty"`
}

type APIResponseBB struct {
	MessageBody []model.BlogPost `json:"data,omitempty"`
}

func TestBlogHandlerImpl_handleBlogPostGet(t *testing.T) {

	// Set up
	sv := service.NewMockBlogService(t)
	w := httptest.NewRecorder()
	router := gin.New()
	b := NewBlogHandler(sv)
	router.GET("/api/v1/posts/:id", b.handleBlogPostGet)

	// Test expected request
	// if you can figure out mock.AnythingOfType(context.Context) do call me
	sv.EXPECT().GetByID(mock.Anything, 1, true).Return(FakeBlogPost, nil)
	req, _ := http.NewRequest("GET", "/api/v1/posts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var retJSON APIResponseB
	json.Unmarshal(w.Body.Bytes(), &retJSON)
	assert.Equal(t, FakeBlogPost.Title, retJSON.MessageBody.Title)
	assert.Equal(t, FakeBlogPost.Content, retJSON.MessageBody.Content)
	assert.Equal(t, FakeBlogPost.Tags, retJSON.MessageBody.Tags)

	// Test Bad Request (0 should count as a bad request)
	req, _ = http.NewRequest("GET", "/api/v1/posts/0", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test unexpected error
	sv.EXPECT().GetByID(mock.Anything, 2, true).Return(nil, errors.New("UnexpectedError"))
	req, _ = http.NewRequest("GET", "/api/v1/posts/2", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Test NotFoundError
	sv.EXPECT().GetByID(mock.Anything, 999, true).Return(nil, &repository.NotFoundError{
		PostID: 999,
	})
	req, _ = http.NewRequest("GET", "/api/v1/posts/999", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBlogHandlerImpl_handleBlogPostPost(t *testing.T) {

	// Set up
	sv := service.NewMockBlogService(t)
	w := httptest.NewRecorder()
	router := gin.New()
	b := NewBlogHandler(sv)
	router.POST("/api/v1/posts", b.handleBlogPostPost)

	// Test expected response
	sv.EXPECT().Post(mock.Anything, FakeBlogPost).Return(FakeBlogPost, nil)

	blogJSON, err := json.Marshal(FakeBlogPost)
	if err != nil {
		t.Fatal(err)
	}
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/posts",
		strings.NewReader(string(blogJSON)),
	)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var retJSON APIResponseB
	json.Unmarshal(w.Body.Bytes(), &retJSON)
	assert.Equal(t, FakeBlogPost.Title, retJSON.MessageBody.Title)
	assert.Equal(t, FakeBlogPost.Content, retJSON.MessageBody.Content)
	assert.Equal(t, FakeBlogPost.Tags, retJSON.MessageBody.Tags)

	// Test 400 Malformed request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(
		"POST",
		"/api/v1/posts",
		strings.NewReader("Bad request"),
	)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test unexpected error
	sv.EXPECT().Post(mock.Anything, FakeBlogPost2).Return(nil, errors.New("UnexpectedError"))
	w = httptest.NewRecorder()
	blogJSON2, _ := json.Marshal(FakeBlogPost2)
	req, _ = http.NewRequest(
		"POST",
		"/api/v1/posts",
		strings.NewReader(string(blogJSON2)),
	)

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestBlogHandlerImpl_handleBlogPostGetAll(t *testing.T) {

	// Set up
	sv := service.NewMockBlogService(t)
	w := httptest.NewRecorder()
	router := gin.New()
	b := NewBlogHandler(sv)
	router.GET("/api/v1/posts", b.handleBlogPostGetAll)

	//Test expected with no query term
	sv.EXPECT().GetAll(mock.Anything, "NONE", true).Return([]*model.BlogPost{FakeBlogPost, FakeBlogPost2}, nil)
	req, _ := http.NewRequest("GET", "/api/v1/posts", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var retJSON APIResponseBB

	json.Unmarshal(w.Body.Bytes(), &retJSON)
	assert.Equal(t, retJSON.MessageBody, 
		[]model.BlogPost{*FakeBlogPost, *FakeBlogPost2})


	//Test with query term
	//This assures that the term query param does get passed to the service
	sv.EXPECT().GetAll(mock.Anything, "TERM", true).Return([]*model.BlogPost{FakeBlogPost, FakeBlogPost2}, nil)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/posts?term=TERM", nil)
		
	router.ServeHTTP(w, req)

	//Test Internal Server Error
	sv.EXPECT().GetAll(mock.Anything, "ERROR", true).Return(nil, errors.New("Boom"))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/posts?term=ERROR", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestBlogHandlerImpl_handleBlogPostDelete(t *testing.T) {
	// Set up
	sv := service.NewMockBlogService(t)
	w := httptest.NewRecorder()
	router := gin.New()
	b := NewBlogHandler(sv)
	router.DELETE("/api/v1/posts/:id", b.handleBlogPostDelete)

	// Test no ID given
	req, _ := http.NewRequest("DELETE", "/api/v1/posts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test malformed ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/api/v1/posts/asdf", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)


	// Test Expected
	w = httptest.NewRecorder()
	sv.EXPECT().Delete(mock.Anything, 1).Return(nil)
	req, _ = http.NewRequest("DELETE", "/api/v1/posts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Test not found
	w = httptest.NewRecorder()
	sv.EXPECT().Delete(mock.Anything, 999).Return(&repository.NotFoundError{PostID: 999})

	req, _ = http.NewRequest("DELETE", "/api/v1/posts/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test internal server error
	w = httptest.NewRecorder()
	sv.EXPECT().Delete(mock.Anything, 888).Return(errors.New("Boom"))
	req, _ = http.NewRequest("DELETE", "/api/v1/posts/888", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)


}

func TestBlogHandlerImpl_handleBlogPostUpdate(t *testing.T) {

	// Set up
	sv := service.NewMockBlogService(t)
	w := httptest.NewRecorder()
	router := gin.New()
	b := NewBlogHandler(sv)
	router.PUT("/api/v1/posts/:id", b.handleBlogPostUpdate)

	// Test Expected
	sv.EXPECT().Update(mock.Anything, 1, FakeBlogPost).Return(nil)
	reqJSON, _ := json.Marshal(FakeBlogPost)
	req, _ := http.NewRequest("PUT", "/api/v1/posts/1", strings.NewReader(string(reqJSON)))
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test malformed ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/posts/asdf", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test no ID given
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/posts", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	// Test Blog Malformed
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/posts/asdf", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test Bad Blog Input does not marshal
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/posts/1", strings.NewReader("BADREQUEST"))
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test NotFoundError
	reqJSON2, _ := json.Marshal(FakeBlogPost2)
	sv.EXPECT().Update(mock.Anything, 2, FakeBlogPost2).Return(&repository.NotFoundError{})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/posts/2", strings.NewReader(string(reqJSON2)))
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	// Test Internal Server Error
	sv.EXPECT().Update(mock.Anything, 500, FakeBlogPost2).Return(errors.New("UnexpectedError"))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/posts/500", strings.NewReader(string(reqJSON2)))
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}
