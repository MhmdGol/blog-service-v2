package main

import (
	"blog-service/internal/controller"
	"blog-service/internal/repository"
	"blog-service/internal/service"
)

func main() {
	repo := repository.New()
	srv := service.New(repo)
	ctrl := controller.New(srv)

	ctrl.Start()
}
