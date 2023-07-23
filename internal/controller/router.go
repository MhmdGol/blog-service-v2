package controller

import (
	"blog-service/internal/controller/authentication"
	"blog-service/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type httpApi struct {
	router *mux.Router
	app    service.App
}

func New(app service.App) *httpApi {
	a := &httpApi{
		router: mux.NewRouter().StrictSlash(true),
		app:    app,
	}

	a.SetupRoutes()
	return a
}

func (h *httpApi) SetupRoutes() {
	h.router.HandleFunc("/login", authentication.LoginHandler).Methods("POST")

	h.router.HandleFunc("/post/create", h.createNewPost).Methods("POST")
	h.router.HandleFunc("/post/read", h.getAllPosts).Methods("GET")
	h.router.HandleFunc("/post/read/{id}", h.getPagePosts).Methods("GET")
	h.router.HandleFunc("/post/update/{id}", h.updatePost).Methods("PUT")
	h.router.HandleFunc("/post/delete/{id}", h.deletePost).Methods("DELETE")

	h.router.HandleFunc("/category/create", h.createNewCategory).Methods("POST")
	h.router.HandleFunc("/category/read", h.getAllCategories).Methods("GET")
	h.router.HandleFunc("/category/update/{id}", h.updateCategory).Methods("PUT")
	h.router.HandleFunc("/category/delete/{id}", h.deleteCategory).Methods("DELETE")
}

func (h *httpApi) Start() {
	// the port must be replaced with config
	log.Fatal(http.ListenAndServe(":8080", h.router))
}
