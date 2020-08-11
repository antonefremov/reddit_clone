package main

import (
	"context"
	"database/sql"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/handlers"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/middleware"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/posts"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/session"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/user"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func main() {

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// MySql DB
	// основные настройки к базе
	dsn := "root:love1234@tcp(localhost:3306)/golang2?"
	// указываем кодировку
	dsn += "&charset=utf8"
	// отказываемся от prapared statements
	// параметры подставляются сразу
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	// -----

	// Mongo DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Errorf("Can't connect to mongodb. %s", err.Error())
		return
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Errorf("MongoDB ping error. %s", err.Error())
		return
	}
	defer func(c *mongo.Client) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := c.Disconnect(ctx)
		if err != nil {
			logger.Error("Can't disconnect from mongodb. %s", err.Error())
		}
	}(client)
	coll := client.Database("asperitas").Collection("posts")
	postsCollection := &posts.MongoCollection{
		Сoll: coll,
	}

	sm := session.NewSessionsManager(db)

	usersRepo := user.NewRepo(db)
	postsRepo := posts.NewRepo(postsCollection)

	usersHandler := &handlers.UsersHandler{
		Logger:    logger,
		UsersRepo: usersRepo,
		Sessions:  sm,
	}

	postsHandler := &handlers.PostsHandler{
		Logger:    logger,
		UsersRepo: usersRepo,
		PostsRepo: postsRepo,
	}

	r := mux.NewRouter()

	// ar := mux.NewRouter()
	// authRoute := middleware.Auth(sm, logger, r)

	postsRouter := r.PathPrefix("/api/posts").Subrouter()
	postsRouter.HandleFunc("/", postsHandler.List).Methods("GET")

	postsRouter.HandleFunc("/{category}", postsHandler.GetListByCat).Methods("GET")

	addChain := middleware.Chain(postsHandler.Add, middleware.AuthorizedUserMiddleware(sm, logger))
	postsRouter.HandleFunc("", addChain).Methods("POST")

	// postsByUserRouter := r.PathPrefix("/api/user/{user_login}").Subrouter()

	postRouter := r.PathPrefix("/api/post").Subrouter()
	postRouter.HandleFunc("/{id}", postsHandler.GetPostByID).Methods("GET")

	deletePostChain := middleware.Chain(postsHandler.Delete, middleware.AuthorizedUserMiddleware(sm, logger))
	postRouter.HandleFunc("/{id}", deletePostChain).Methods("DELETE")

	addCommentChain := middleware.Chain(postsHandler.AddComment, middleware.AuthorizedUserMiddleware(sm, logger))
	postRouter.HandleFunc("/{id}", addCommentChain).Methods("POST")

	deleteCommentChain := middleware.Chain(postsHandler.DeleteComment, middleware.AuthorizedUserMiddleware(sm, logger))
	postRouter.HandleFunc("/{id}/{commentId}", deleteCommentChain).Methods("DELETE")

	upvote := middleware.Chain(postsHandler.Upvote, middleware.AuthorizedUserMiddleware(sm, logger))
	postRouter.HandleFunc("/{id}/upvote", upvote).Methods("GET")

	unvote := middleware.Chain(postsHandler.Unvote, middleware.AuthorizedUserMiddleware(sm, logger))
	postRouter.HandleFunc("/{id}/unvote", unvote).Methods("GET")

	downvote := middleware.Chain(postsHandler.Downvote, middleware.AuthorizedUserMiddleware(sm, logger))
	postRouter.HandleFunc("/{id}/downvote", downvote).Methods("GET")

	usersRouter := r.PathPrefix("/api").Subrouter()
	usersRouter.HandleFunc("/login", usersHandler.Login).Methods("POST")
	usersRouter.HandleFunc("/logout", usersHandler.Logout).Methods("POST")
	usersRouter.HandleFunc("/register", usersHandler.Register).Methods("POST")

	usersRouter.HandleFunc("/user/{user_login}", postsHandler.GetListByAuthor).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./../../template")))

	addr := ":8080"
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	http.ListenAndServe(addr, r)
}
