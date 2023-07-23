package repository

import (
	"context"

	"gorm.io/gorm"
)

type Store interface {
	CreatePost(ctx context.Context, title, text string, categories []string) error
	GetAllPosts(ctx context.Context) ([]Post, error)
	GetPagePosts(ctx context.Context) ([]Post, error)
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
