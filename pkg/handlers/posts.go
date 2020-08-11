package handlers

import (
	"encoding/json"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/posts"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/session"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/utils"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// PostsRepoInterface represents methods available for the PostsRepo
type PostsRepoInterface interface {
	All() ([]*posts.Post, error)
	ListByCategory(string) ([]*posts.Post, error)
	Get(string) (*posts.Post, error)
	GetByAuthor(string) ([]*posts.Post, error)
	Add(*posts.Post) (*posts.Post, error)
	AddComment(string, *posts.Comment) (*posts.Post, error)
	DeleteComment(string, string) (*posts.Post, error)
	Delete(string, string) error
	Vote(string, posts.Vote) (*posts.Post, error)
	Unvote(string, string) (*posts.Post, error)
}

// PostsHandler is a hook to work with incoming requests for the Posts collection
type PostsHandler struct {
	PostsRepo PostsRepoInterface // *posts.Repo
	UsersRepo UsersRepoInterface // *user.Repo
	Logger    *zap.SugaredLogger
}

// List returns all Post objects (full list)
func (h *PostsHandler) List(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostsRepo.All()
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(posts)
	w.Write(result)
}

// GetListByCat returns a list of Post objects filtered by a category
func (h *PostsHandler) GetListByCat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	posts, err := h.PostsRepo.ListByCategory(category)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(posts)
	w.Write(result)
}

// GetListByAuthor returns a list of Post objects filtered by a category
func (h *PostsHandler) GetListByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userLogin := vars["user_login"]

	posts, err := h.PostsRepo.GetByAuthor(userLogin)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	result, _ := json.Marshal(posts)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// GetPostByID returns a Post objects by its Id
func (h *PostsHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post, err := h.PostsRepo.Get(id)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(post)
	w.Write(result)
}

// Add creates a new Post in the repository
func (h *PostsHandler) Add(w http.ResponseWriter, r *http.Request) {

	newPost := &posts.Post{}
	err := json.NewDecoder(r.Body).Decode(newPost)
	if err != nil {
		h.Logger.Errorf(`BadRequest. %s`, err.Error())
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusBadRequest)
		return
	}

	sess, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	author, err := h.UsersRepo.GetByID(sess.UserID)
	if err != nil {
		h.Logger.Errorf(`InternalServerError. Could not find user. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}
	author.PasswordHash = "" // a security measure :)
	newPost.Author = *author

	createdPost, err := h.PostsRepo.Add(newPost)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(createdPost)
	w.Write(result)
}

// Delete removes an existing Post in the repository
func (h *PostsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sess, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	err = h.PostsRepo.Delete(id, sess.UserID)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	result := utils.GetJSONMessageAsString("success")
	io.WriteString(w, result)
}

// Upvote adds up a user's vote to a Post
func (h *PostsHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sess, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	vote := &posts.Vote{User: sess.UserID, Vote: 1}

	post, err := h.PostsRepo.Vote(id, *vote)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(post)
	w.Write(result)
}

// Unvote removes previously added user's vote on a Post
func (h *PostsHandler) Unvote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sess, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	post, err := h.PostsRepo.Unvote(id, sess.UserID)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(post)
	w.Write(result)
}

// Downvote removes a user's vote from a Post
func (h *PostsHandler) Downvote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sess, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}
	vote := &posts.Vote{User: sess.UserID, Vote: -1}

	post, err := h.PostsRepo.Vote(id, *vote)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(post)
	w.Write(result)
}

// AddComment adds a new comment to a Post
func (h *PostsHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	newComment := &posts.Comment{}
	payload := &posts.NetworkComment{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newComment.Body = payload.Comment

	sess, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	author, err := h.UsersRepo.GetByID(sess.UserID)
	if err != nil {
		h.Logger.Errorf(`InternalServerError. Could not find user. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}
	author.PasswordHash = "" // a security measure :)
	newComment.Author = *author

	post, err := h.PostsRepo.AddComment(postID, newComment)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(post)
	w.Write(result)
}

// DeleteComment removes an existing comment from a Post
func (h *PostsHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	commentID := vars["commentId"]

	_, err := session.SessionFromContext(r.Context())
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	post, err := h.PostsRepo.DeleteComment(postID, commentID)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(post)
	w.Write(result)
}
