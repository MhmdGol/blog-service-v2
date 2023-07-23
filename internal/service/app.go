package service

import (
	"blog-service/internal/repository"
	"context"
)

type App interface {
	CreatePost(ctx context.Context, title, text string, categories []string) error
	GetAllPosts(ctx context.Context) ([]repository.Post, error)
	GetPagePosts(ctx context.Context, pageNumber, pageSize int) ([]repository.Post, error)
	UpdatePost(ctx context.Context, postID int, title, text string, categories []string) error
	DeletePost(ctx context.Context, postID int) error

	CreateCategory(ctx context.Context, name string) error
	GetAllCategories(ctx context.Context) ([]repository.Category, error)
	UpdateCategory(ctx context.Context, categoryID int, name string) error
	DeleteCategory(ctx context.Context, categoryID int) error
}

type app struct {
	store repository.Store
}

func New(store repository.Store) *app {
	return &app{
		store: store,
	}
}
