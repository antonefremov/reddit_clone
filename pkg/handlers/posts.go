package handlers

import (
	"encoding/json"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/ljwt"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/posts"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/users"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/utils"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// PostsHandler is a hook to work with incoming requests for the Posts collection
type PostsHandler struct {
	PostsRepo *posts.Repo
	// Logger    *zap.SugaredLogger
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
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	tokenUser, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
		return
	}

	author := &users.User{
		Username: tokenUser.Username,
		ID:       tokenUser.ID,
	}
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

	ctx := r.Context()
	tokenUser, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
		return
	}

	err := h.PostsRepo.Delete(id, tokenUser.ID)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	result := utils.GetJSONMessageAsString("success")
	// http.Redirect(w, r, "/", 302)
	io.WriteString(w, result)
}

// Upvote adds up a user's vote to a Post
func (h *PostsHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	tokenUser, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
		return
	}
	vote := &posts.Vote{User: tokenUser.ID, Vote: 1}

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

	ctx := r.Context()
	tokenUser, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
		return
	}

	post, err := h.PostsRepo.Unvote(id, tokenUser.ID)
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

	ctx := r.Context()
	tokenUser, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
		return
	}
	vote := &posts.Vote{User: tokenUser.ID, Vote: -1}

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

	ctx := r.Context()
	tokenUser, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
		return
	}

	author := new(users.User)
	author.Username = tokenUser.Username
	author.Admin = false
	author.ID = tokenUser.ID
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

	ctx := r.Context()
	_, ok := ctx.Value(utils.CurrentUserKey).(*ljwt.TokenUser)
	if !ok {
		jsonMessage := utils.GetJSONMessageAsString("Authentication failed")
		http.Error(w, jsonMessage, http.StatusForbidden)
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
