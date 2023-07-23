package repository

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store interface {
	CreatePost(ctx context.Context, title, text string, categories []string) error
	GetAllPosts(ctx context.Context) ([]Post, error)
	GetPagePosts(ctx context.Context, pageNumber, pageSize int) ([]Post, error)
	UpdatePost(ctx context.Context, postID int, title, text string, categories []string) error
	DeletePost(ctx context.Context, postID int) error

	CreateCategory(ctx context.Context, name string) error
	GetAllCategories(ctx context.Context) ([]Category, error)
	UpdateCategory(ctx context.Context, categoryID int, name string) error
	DeleteCategory(ctx context.Context, categoryID int) error
}

type store struct {
	db *gorm.DB
}

func New() *store {
	// must be replaced by viper configs
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=blog-service sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&Category{}, &Post{})
	if err != nil {
		panic(err)
	}

	return &store{
		db: db,
	}
}
