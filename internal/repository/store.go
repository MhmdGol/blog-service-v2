package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store interface {
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
