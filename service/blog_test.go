package service

import (
	"context"
	"testing"

	"github.com/go-playground/assert"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
)

func TestBlogStandard_GetAllNoContent(t *testing.T) {

	ctx := context.Background()
	mockrepository := repository.NewMockBlogRepository(t)



	sv := NewBlogService(mockrepository)
	mockrepository.EXPECT().GetAll(ctx, SummaryCard).Return([]*model.BlogPost{}, nil)
	x, err := sv.GetAllNoContent(ctx)
	if err != nil {
		t.Fatalf("Error in Test blogservice.GetAll: %s", err)
	}

	// Assert that it does return blogposts
	assert.Equal(t, x, []*model.BlogPost{})
}
