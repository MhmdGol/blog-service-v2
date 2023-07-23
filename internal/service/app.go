package service

import (
	"blog-service/internal/repository"
)

type App interface {
	CreatePost(title, text string, categories []string) error
	GetAllPosts() ([]Post, error)
	GetPagePosts(pageNumber, pageSize int) ([]Post, error)
	UpdatePost(postID int, title, text string, categories []string) error
	DeletePost(postID int) error

	CreateCategory(name string) error
	GetAllCategories() ([]Category, error)
	UpdateCategory(categoryID int, name string) error
	DeleteCategory(categoryID int) error
}

type app struct {
	store repository.Store
}

func New(store repository.Store) *app {
	return &app{
		store: store,
	}
}
