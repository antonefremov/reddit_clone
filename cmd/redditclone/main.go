package main

import (
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/handlers"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/middleware"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/posts"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// sm := session.NewSessionsMem()
	usersRepo := users.NewUsersRepo()
	postsRepo := posts.NewRepo()

	usersHandler := &handlers.UsersHandler{
		// Logger:    logger,
		UsersRepo: usersRepo,
		// Sessions:  sm,
	}

	postsHandler := &handlers.PostsHandler{
		// Logger:    logger,
		PostsRepo: postsRepo,
	}

	r := mux.NewRouter()

	postsRouter := r.PathPrefix("/api/posts").Subrouter()
	postsRouter.HandleFunc("/", postsHandler.List).Methods("GET").
		Schemes("http").
		Host("localhost")
	postsRouter.HandleFunc("/{category}", postsHandler.GetListByCat).Methods("GET").
		Schemes("http").
		Host("localhost")

	addChain := middleware.Chain(postsHandler.Add, middleware.AuthorizedUserMiddleware())
	postsRouter.HandleFunc("", addChain).Methods("POST").
		Schemes("http").
		Host("localhost")

	postRouter := r.PathPrefix("/api/post").Subrouter()
	postRouter.HandleFunc("/{id}", postsHandler.GetPostByID).Methods("GET").
		Schemes("http").
		Host("localhost")

	deletePostChain := middleware.Chain(postsHandler.Delete, middleware.AuthorizedUserMiddleware())
	postRouter.HandleFunc("/{id}", deletePostChain).Methods("DELETE").
		Schemes("http").
		Host("localhost")

	addCommentChain := middleware.Chain(postsHandler.AddComment, middleware.AuthorizedUserMiddleware())
	postRouter.HandleFunc("/{id}", addCommentChain).Methods("POST").
		Schemes("http").
		Host("localhost")

	deleteCommentChain := middleware.Chain(postsHandler.DeleteComment, middleware.AuthorizedUserMiddleware())
	postRouter.HandleFunc("/{id}/{commentId}", deleteCommentChain).Methods("DELETE").
		Schemes("http").
		Host("localhost")

	upvote := middleware.Chain(postsHandler.Upvote, middleware.AuthorizedUserMiddleware())
	postRouter.HandleFunc("/{id}/upvote", upvote).Methods("GET").
		Schemes("http").
		Host("localhost")

	unvote := middleware.Chain(postsHandler.Unvote, middleware.AuthorizedUserMiddleware())
	postRouter.HandleFunc("/{id}/unvote", unvote).Methods("GET").
		Schemes("http").
		Host("localhost")

	downvote := middleware.Chain(postsHandler.Downvote, middleware.AuthorizedUserMiddleware())
	postRouter.HandleFunc("/{id}/downvote", downvote).Methods("GET").
		Schemes("http").
		Host("localhost")

	usersRouter := r.PathPrefix("/api").Subrouter()
	usersRouter.HandleFunc("/login", usersHandler.Login).Methods("POST").
		Schemes("http").
		Host("localhost")
	usersRouter.HandleFunc("/logout", usersHandler.Logout).Methods("POST").
		Schemes("http").
		Host("localhost")
	usersRouter.HandleFunc("/register", usersHandler.Register).Methods("POST").
		Schemes("http").
		Host("localhost")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./../../template")))

	log.Println("Starting server at :8080")
	http.ListenAndServe(":8080", r)

}
