package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/iqunlim/easyblog/model"
	"gorm.io/gorm"
)


type BlogRepository interface {
	Post(ctx context.Context, blog *model.BlogPost) error
	Update(ctx context.Context,id int, updateFn func(*model.BlogPost) (bool, error)) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int, fields []string) (*model.BlogPost, error)
	GetByFilter(ctx context.Context, queryparam string, fields []string) ([]*model.BlogPost, error)
	GetAll(ctx context.Context, fields []string) ([]*model.BlogPost, error)
}


type BlogRepositoryImpl struct {
	DB *gorm.DB
}

func NewBlogRepository(DB *gorm.DB) BlogRepository {
	return &BlogRepositoryImpl{
		DB: DB,
	}
}

func (b *BlogRepositoryImpl) Post(ctx context.Context, blog *model.BlogPost) error { 

	return b.DB.Create(blog).Error

}
func (b *BlogRepositoryImpl) Update(ctx context.Context,id int, updateFn func(updatingPost *model.BlogPost) (bool, error)) error { 


	var post model.BlogPost
	if err := b.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &NotFoundError{
				PostID: id,
			}
		}
		return err
	}

	updated, err := updateFn(&post)
	if err != nil {
		return err
	}
	if !updated {
		return nil
	}

	if err = b.DB.Save(&post).Error; err != nil {
		return err
	}

	return nil
}
func (b *BlogRepositoryImpl) Delete(ctx context.Context,id int) error { 
	

	res := b.DB.Delete(&model.BlogPost{}, id)
	err := res.Error

	if err != nil {
		return err
	}
	if res.RowsAffected != 1 {
		return &NotFoundError{
			PostID: id,
		}
	}
	return nil
}

func (b *BlogRepositoryImpl) GetByID(ctx context.Context,id int, fields []string) (*model.BlogPost, error) { 

	var post *model.BlogPost
	if err := b.DB.First(&post, id).Select(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &NotFoundError{
				PostID: id,
			}
		}
		return nil, err
	}
	return post, nil 
}


func (b *BlogRepositoryImpl) GetByFilter(ctx context.Context, queryparam string, fields []string) ([]*model.BlogPost, error) { 

	var posts []*model.BlogPost
	// SELECT * FROM posts WHERE title LIKE %queryparam% OR tags LIKE %queryparam% OR Content LIKE %queryparam%
	if err := b.DB.Where("tags LIKE ?", queryparam).
	//Or("title LIKE ?", queryparam).
	//Or("content LIKE ?", queryparam).
	Select(fields).
	Find(&posts).Order("created_at").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &NotFoundError{}
		}
		return nil, err
	}
	return posts, nil
}

func (b *BlogRepositoryImpl) GetAll(ctx context.Context, fields []string) ([]*model.BlogPost, error) {

	var posts []*model.BlogPost
	if err := b.DB.Select(fields).Find(&posts).Order("created_at").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &NotFoundError{}
		}
		return nil, err
	}
	return posts, nil
}

type NotFoundError struct {
	PostID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Post ID: %d Not Found", e.PostID)
}
