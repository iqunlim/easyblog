package service

import (
	"context"
	"log"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
)

var (
	AllBlogFields = []string{
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Title", "Content", "Category", "Tags",
	}

	TitleOnly = []string{
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Title",
	}
)



type BlogService interface {
	Post(ctx context.Context, blog *model.BlogPost) (*model.BlogPost, error)
	Update(ctx context.Context, id int, blog *model.BlogPost) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int, htmlformat bool) (*model.BlogPost, error)
	GetAll(ctx context.Context, params string, htmlformat bool) ([]*model.BlogPost, error)
	GetAllNoContent(ctx context.Context) ([]*model.BlogPost, error)
}

type BlogStandard struct {
	repository repository.BlogRepository
}

func NewBlogService(repository repository.BlogRepository) BlogService {
	return &BlogStandard{
		repository: repository,
	}
}

func (b *BlogStandard) Post(ctx context.Context, blog *model.BlogPost) (*model.BlogPost, error) {
	// Do sanitizing of inputs here
	// Do logging here
	// Do verification here

	return blog, b.repository.Post(ctx, blog)
}


func (b *BlogStandard) Update(ctx context.Context, id int, blog *model.BlogPost) error {

	updateFn := func(updatingPost *model.BlogPost) (bool, error) {

		updated := false
		if updatingPost.Title != blog.Title && blog.Title != "" {
			updatingPost.Title = blog.Title
			log.Println("Title updated")
			updated = true
		}

		if updatingPost.Content != blog.Content && blog.Content != "" {
			updatingPost.Content = blog.Content 
			log.Println("Content updated")
			updated = true
		}

		if updatingPost.Category != blog.Category && blog.Category != "" {
			updatingPost.Category = blog.Category
			log.Println("Category updated")
			updated = true
		}
		if blog.Tags != nil {
			updatingPost.Tags = blog.Tags
			updated = true
		}
		log.Println("Updated: ", updated)
		return updated, nil
	}
	return b.repository.Update(ctx, id, updateFn)
}

func (b *BlogStandard) Delete(ctx context.Context, id int) error {
	return b.repository.Delete(ctx, id)
}

func (b *BlogStandard) GetByID(ctx context.Context, id int, htmlformat bool) (*model.BlogPost, error) {
	res, err := b.repository.GetByID(ctx, id, AllBlogFields)
	if err != nil {
		return nil, err
	}
	if htmlformat {
		FormatBlogForHTML(res)
	}
	return res, err
}

func (b *BlogStandard) GetAll(ctx context.Context, params string, htmlformat bool) ([]*model.BlogPost, error) {

	// Eh...

	var p []*model.BlogPost
	var err error
	if (params == "NONE" || params == "") {

		p, err = b.repository.GetAll(ctx, AllBlogFields)
		if err != nil {
			return nil, err
		}
	} else {
		params = "%" + params + "%"
		p, err = b.repository.GetByFilter(ctx, params, AllBlogFields)
		if err != nil {
			return nil, err
		}
	}
	if htmlformat {
		for _, post := range p {
			FormatBlogForHTML(post)
		}
	}
	return p, nil
}


func (b *BlogStandard) GetAllNoContent(ctx context.Context) ([]*model.BlogPost, error) {
	return b.repository.GetAll(ctx, TitleOnly)
}


//TODO: Create some sort of Format struct that formatting functions can implement the interface of
func FormatBlogForHTML(post *model.BlogPost) {

	md := []byte(post.Content)
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	post.Content = string(markdown.Render(doc, renderer))
	
}


/*
func test() {
	x := model.BlogPost{}
	Format(&x, FormatBlogForHTML)
}
*/
/*
func Format(post *model.BlogPost, formats ...Formatter) {

	for _, function := range formats {
		function(post)
	}
}

// Function takes in a model.BlogPost and formats it.
// Example: FormatBlogForHTML
type Formatter func(post *model.BlogPost) 
*/
