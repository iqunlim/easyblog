package repository

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var (
	AllBlogFields = []string{
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Title", "Content", "Summary", "ImageUrl", "Tags",
	}

	TitleOnly = []string{
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Title", "Tags",
	}

	SummaryCard = []string{
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Title", "Tags", "Summary", "ImageUrl",
	}
)


func NewMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	// The dumbest. There is no way to stop gorm from checking sqlite version
	mock.ExpectQuery("select sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.46.1"))
	if err != nil {
		log.Fatalf("Unexpected Error happened on mock DB creation: %s", err)
	}
	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{Conn: db}), &gorm.Config{SkipDefaultTransaction: true, DisableAutomaticPing: true})
	if err != nil {
		log.Fatalf("Unexpected Error happend on gorm sqlite connection creation: %s", err)
	}
	return gormDB, mock
}

func TestBlogRepositoryImpl_GetAll(t *testing.T) {
	gdb, mock := NewMockDB(t)

	blogrepo := NewBlogRepository(gdb)

	mock.ExpectQuery("SELECT `id`,`created_at`,`updated_at`,`deleted_at`,`title`,`content`,`summary`,`image_url`,`tags` FROM `blog_posts` WHERE `blog_posts`.`deleted_at` IS NULL").
	WillReturnRows(sqlmock.NewRows(AllBlogFields))
	res, err := blogrepo.GetAll(context.Background(), AllBlogFields)
	if err != nil {
		t.Fatalf("Error in TestingBlogRepository_GetAll, %s", err)
	}
	log.Println(res)

}
