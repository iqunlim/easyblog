package controller

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/iqunlim/easyblog/config"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"github.com/iqunlim/easyblog/service"
	mock "github.com/stretchr/testify/mock"
)



func setupWeb(t *testing.T) (*gin.Engine, WebHandler, *service.MockBlogService) {
	//t.Helper()
	sv := service.NewMockBlogService(t)
	router := gin.New()
	web := NewWebHandler(sv)
	router.SetFuncMap(template.FuncMap{
		"formatAsHTML":       formatAsHTML,
		"formatTimeToString": formatTimeToString,
	})

	// Static file handling for the webserver portion
	router.LoadHTMLGlob("../static/template/*.html")
	router.Static("../static", "./static")
	return router, web, sv
}

func bodyHasFragments(t *testing.T, body string, fragments []string) {
	t.Helper()
	for _, fragment := range fragments {
		if !strings.Contains(body, fragment) {
			t.Fatalf("expected body to contain '%s', got %s",
				fragment,
				body)
		}
	}
}
func bodyNotHasFragments(t *testing.T, body string, fragments []string) {
	t.Helper()
	for _, fragment := range fragments {
		if strings.Contains(body, fragment) {
			t.Fatalf("expected body to NOT contain '%s', got %s",
				fragment,
				body)
		}
	}
}
func TestWebHandlerImpl_AdminEditPostHandler(t *testing.T) {

	router, web, sv := setupWeb(t)
	router.GET("/admin/edit", web.AdminEditPostHandler)
	router.GET("/admin/edit/:id", web.AdminEditPostHandler)

	// Test expected with id
	sv.EXPECT().GetByID(mock.Anything, 1, false).Return(FakeBlogPost, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/edit/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	fragments := []string{
		"<title>Edit Post</title>",
	}
	bodyHasFragments(t, w.Body.String(), fragments)

	// Test expected without id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/admin/edit", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	fragments = []string{
		"<title>New Post</title>",
	}
	bodyHasFragments(t, w.Body.String(), fragments)

	// Test unexpeted error from service
	sv.EXPECT().GetByID(mock.Anything, 999, false).Return(nil, errors.New("UnexpectedError"))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/admin/edit/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	fragments = []string{
		"<title>Internal Server Error</title>",
	}

	bodyHasFragments(t, w.Body.String(), fragments)

	// Test 404 Not Found
	sv.EXPECT().GetByID(mock.Anything, 404, false).Return(nil, &repository.NotFoundError{})

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/admin/edit/404", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	fragments = []string{
		"<title>New Post</title>",
	}

	bodyHasFragments(t, w.Body.String(), fragments)
	
	// Test malformed ID

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/admin/edit/asdf", nil)

	router.ServeHTTP(w, req)

	fragments = []string{
		"<title>New Post</title>",
	}

	assert.Equal(t, http.StatusOK, w.Code)
	bodyHasFragments(t, w.Body.String(), fragments)

	// Test ID = 0

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/admin/edit/0", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fragments = []string{
		"<title>New Post</title>",
	}

	bodyHasFragments(t, w.Body.String(), fragments)
}

func TestWebHandlerImpl_AdminDashboardHandler(t *testing.T) {

	router, web, sv := setupWeb(t)
	router.GET("/admin", web.AdminDashboardHandler)

	// Test Expected
	ret := FakeBlogPost
	ret.ID = 1

	ret2 := FakeBlogPost2
	ret2.ID = 2
	sv.EXPECT().GetAllNoContent(mock.Anything).Return([]*model.BlogPost{ret, ret2}, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fragments := []string{
		"Test <a class=\"edit\" href=\"/admin/edit/1\">",
		"Test2 <a class=\"edit\" href=\"/admin/edit/2\">",
	}

	bodyHasFragments(t, w.Body.String(), fragments)

	// Test No Posts at all
}

func TestWebHandlerImpl_AdminDashboardHandler_UNEXPECTED(t *testing.T) {

	// Test Unexpected Error
	router, web, sv := setupWeb(t)
	router.GET("/admin", web.AdminDashboardHandler)
	sv.EXPECT().GetAllNoContent(mock.Anything).Return(nil, errors.New("UnexpectedError"))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	fragments := []string{
		"Internal Server Error",
	}

	bodyHasFragments(t, w.Body.String(), fragments)
}
func TestWebHandlerImpl_AdminDashboardHandler_NOPOSTS(t *testing.T) {

	// Test Not Found Error
	router, web, sv := setupWeb(t)
	router.GET("/admin", web.AdminDashboardHandler)
	sv.EXPECT().GetAllNoContent(mock.Anything).Return(nil, &repository.NotFoundError{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	// Fragment of the edit class, should be none if no posts
	fragments := []string{
		"<a class=\"edit\"",
	}

	bodyNotHasFragments(t, w.Body.String(), fragments)
}

func TestWebHandlerImpl_PostsByIDWebHandler(t *testing.T) {
	router, web, sv := setupWeb(t)
	router.GET("/posts/:id", web.PostsByIDWebHandler)

	// Test Expected
	sv.EXPECT().GetByID(mock.Anything, 1, true).Return(FakeBlogPost, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	fragments := []string{
		"<title>Test</title>",
	}

	bodyHasFragments(t, w.Body.String(), fragments)

	// Test Bad ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts/asf", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	fragments = []string{
		"Malformed Request",
	}

	bodyHasFragments(t, w.Body.String(), fragments)

	// Test 0 ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	fragments = []string{
		"Malformed Request",
	}

	bodyHasFragments(t, w.Body.String(), fragments)
	// Test Not Found
	sv.EXPECT().GetByID(mock.Anything, 999, true).Return(nil, &repository.NotFoundError{})

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts/999", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)

	// Renders the 302 html before it goes through
	fragments = []string{
		"<a href=\"/posts\"",
	}

	bodyHasFragments(t, w.Body.String(), fragments)

	// Test Unexpected Error
	sv.EXPECT().GetByID(mock.Anything, 888, true).Return(nil, errors.New("Boom"))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts/888", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)

	fragments = []string{
		"Post Not Found",
	}

	bodyHasFragments(t, w.Body.String(), fragments)
}

func TestWebHandlerImpl_PostsWebHandler(t *testing.T) {

	router, web, sv := setupWeb(t)
	router.GET("/posts", web.PostsWebHandler)

	// Test expected with no term
	sv.EXPECT().GetAll(mock.Anything, "NONE", true).Return([]*model.BlogPost{FakeBlogPost}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	fragments := []string{
		"<title>All Posts</title>",
	}
	bodyHasFragments(t, w.Body.String(), fragments)
	// Test expected with term
	sv.EXPECT().GetAll(mock.Anything, "TECH", true).Return([]*model.BlogPost{FakeBlogPost2}, nil)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts?term=TECH", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	fragments = []string{
		"<title>All Posts</title>",
	}
	bodyHasFragments(t, w.Body.String(), fragments)

	// Test Unexpected error
	sv.EXPECT().GetAll(mock.Anything, "BOOM", true).Return(nil, errors.New("UnexpectedError"))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts?term=BOOM", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	fragments = []string{
		"Internal Server Error",
	}
	bodyHasFragments(t, w.Body.String(), fragments)
}

func TestWebHandlerImpl_RegisterWebHandlerEnabled(t *testing.T) {

	// TODO handler
	router, web, _ := setupWeb(t)
	router.GET("/register", web.RegisterWebHandler)

	config.RegisterAllowed = "true" // Hard-configuring
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/register", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestWebHandlerImpl_RegisterWebHandlerDisabled(t *testing.T) {

	// TODO handler
	router, web, _ := setupWeb(t)
	router.GET("/register", web.RegisterWebHandler)
	config.RegisterAllowed = "false" // Hard-configuring

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/register", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)

}
func TestWebHandlerImpl_LoginWebHandler(t *testing.T) {

	//TODO: Think about if checking a cookie is okay??? Maybe we 
	// Set something on  the context if authed instead

}
