package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/posts"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/session"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/user"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func TestHandlerGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersRepo := NewMockUsersRepoInterface(ctrl)
	postsRepo := NewMockPostsRepoInterface(ctrl)

	service := PostsHandler{
		PostsRepo: postsRepo,
		UsersRepo: usersRepo,
		Logger:    zap.NewNop().Sugar(),
	}

	login := "login"
	uid := "userid"
	pid := "postid"
	cid := "commentid"

	resultUser := &user.User{
		Username: login, ID: uid,
	}

	resultPost := &posts.Post{
		Score:    0,
		Views:    0,
		Type:     "text",
		Title:    "title",
		Text:     "text",
		Author:   *resultUser,
		Category: "music",
		Votes: []posts.Vote{
			{User: uid, Vote: 1},
		},
		Comments: []posts.Comment{
			{
				Created: time.Now(),
				Author:  *resultUser,
				Body:    "comment_body",
				ID:      cid,
			},
		},
		Created:          time.Now(),
		UpvotePercentage: 100,
		ID:               pid,
	}

	// positive All method
	postsRepo.EXPECT().All().Return([]*posts.Post{resultPost}, nil)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	service.List(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img, _ := json.Marshal([]*posts.Post{resultPost})
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// err All method
	postsRepo.EXPECT().All().Return(nil, errors.New("Database error"))

	req = httptest.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()

	service.List(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`Database error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), string(img))
		return
	}

	//////////////////////////////////////////////////
	// positive ByCategory
	postsRepo.EXPECT().ListByCategory("music").Return([]*posts.Post{resultPost}, nil)

	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"category": "music",
	})
	w = httptest.NewRecorder()

	service.GetListByCat(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img, _ = json.Marshal([]*posts.Post{resultPost})
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// err ByCategory
	postsRepo.EXPECT().ListByCategory("music").Return(nil, errors.New("DB Error"))

	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"category": "music",
	})
	w = httptest.NewRecorder()

	service.GetListByCat(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB Error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	//////////////////////////////////////////
	// positive GetByID
	postsRepo.EXPECT().Get(pid).Return(resultPost, nil)

	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.GetPostByID(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img, _ = json.Marshal(resultPost)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// err GetByID
	postsRepo.EXPECT().Get(pid).Return(nil, errors.New("DB Error"))

	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.GetPostByID(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB Error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// positive GetByAuthor
	postsRepo.EXPECT().GetByAuthor(login).Return([]*posts.Post{resultPost}, nil)

	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"user_login": login,
	})
	w = httptest.NewRecorder()

	service.GetListByAuthor(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img, _ = json.Marshal([]*posts.Post{resultPost})
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// err GetByAuthor
	postsRepo.EXPECT().GetByAuthor(login).Return(nil, errors.New("DB Error"))

	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"user_login": login,
	})
	w = httptest.NewRecorder()

	service.GetListByAuthor(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB Error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}
}

func TestHandlerUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := NewMockUsersRepoInterface(ctrl)
	postsRepo := NewMockPostsRepoInterface(ctrl)

	service := PostsHandler{
		UsersRepo: userRepo,
		PostsRepo: postsRepo,
		Logger:    zap.NewNop().Sugar(),
	}

	login := "login"
	uid := "userid"
	pid := "postid"
	cid := "commentid"
	sid := "sessionid"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uX2lkIjoiaWRkIiwidXNlciI6eyJ1c2VybmFtZSI6ImxvZ2dnZ2dpbiIsImlkIjoiaWRkZGRkZCJ9fQ.AnBam3t75fSHoWx6lFnzp1MBW85ZQKf4ee5SshSiLHk"

	resultUser := &user.User{
		Username: login, ID: uid,
	}

	resultPost := &posts.Post{
		Score:    1,
		Views:    0,
		Type:     "text",
		Title:    "title",
		Text:     "text",
		Author:   *resultUser,
		Category: "programming",
		Votes: []posts.Vote{
			{User: uid, Vote: 1},
		},
		Created:          time.Now().Truncate(0),
		UpvotePercentage: 100,
		ID:               pid,
	}

	// good add
	userRepo.EXPECT().GetByID(uid).Return(resultUser, nil)
	postsRepo.EXPECT().Add(resultPost).Return(resultPost, nil)

	res, _ := json.Marshal(resultPost)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(res))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	w := httptest.NewRecorder()

	service.Add(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img, _ := res, 0
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add bad json body
	req = httptest.NewRequest("POST", "/", bytes.NewReader([]byte{'{'}))
	w = httptest.NewRecorder()

	service.Add(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`{"message":"unexpected EOF"}`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add with bad session
	bts, _ := json.Marshal(resultPost)
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	w = httptest.NewRecorder()

	service.Add(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add with repo err
	userRepo.EXPECT().GetByID(uid).Return(resultUser, nil)
	postsRepo.EXPECT().Add(resultPost).Return(nil, errors.New("DB Error"))

	bts, _ = json.Marshal(resultPost)
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	w = httptest.NewRecorder()

	service.Add(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB Error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add without a proper user found
	userRepo.EXPECT().GetByID(uid).Return(nil, errors.New("No user in the DB"))

	res, _ = json.Marshal(resultPost)
	req = httptest.NewRequest("POST", "/", bytes.NewReader(res))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	w = httptest.NewRecorder()

	service.Add(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	///////////////////////////////////////////////
	// delete post. positive one
	postsRepo.EXPECT().Delete(pid, uid).Return(nil)

	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.Delete(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`{"message":"success"}`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	//delete post. bad id
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	w = httptest.NewRecorder()

	service.Delete(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// delete post. repo err
	postsRepo.EXPECT().Delete(pid, uid).Return(errors.New("Error deleting a Post"))

	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.Delete(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`Error deleting a Post`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	/////////////////////////////////////////////////////////////////////////
	cbody := "commentbody"
	comment := &posts.Comment{
		// Created: time.Now(),
		Author: *resultUser,
		Body:   cbody,
		ID:     "", // cid,
	}
	resultPost.Comments = []posts.Comment{
		*comment,
	}

	// add comment. good
	userRepo.EXPECT().GetByID(uid).Return(resultUser, nil)
	postsRepo.EXPECT().AddComment(pid, comment).Return(resultPost, nil)

	bts, _ = json.Marshal(&posts.NetworkComment{Comment: cbody})
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.AddComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	bts, _ = json.Marshal(resultPost)
	img, _ = bts, 0
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add comment. repo add comment error
	userRepo.EXPECT().GetByID(uid).Return(resultUser, nil)
	postsRepo.EXPECT().AddComment(pid, comment).Return(resultPost, errors.New("DB Error"))

	bts, _ = json.Marshal(&posts.NetworkComment{Comment: cbody})
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.AddComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB Error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add comment. user repo err
	userRepo.EXPECT().GetByID(uid).Return(nil, errors.New(""))

	bts, _ = json.Marshal(&posts.NetworkComment{Comment: cbody})
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.AddComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add comment. sess error
	bts, _ = json.Marshal(&posts.NetworkComment{Comment: cbody})
	req = httptest.NewRequest("POST", "/", bytes.NewReader(bts))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.AddComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// add comment. body err
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.AddComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`EOF`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	/////////////////////////////////////
	resultPost.Comments = nil
	// delete comment. good
	postsRepo.EXPECT().DeleteComment(pid, cid).Return(resultPost, nil)

	req = httptest.NewRequest("POST", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id":        pid,
		"commentId": cid,
	})
	w = httptest.NewRecorder()

	service.DeleteComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	bts, _ = json.Marshal(resultPost)
	img = bts
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// delete comment. repo delete err
	postsRepo.EXPECT().DeleteComment(pid, cid).Return(nil, errors.New("DB error"))

	req = httptest.NewRequest("POST", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id":        pid,
		"commentId": cid,
	})
	w = httptest.NewRecorder()

	service.DeleteComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// delete comment. no sess
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id":        pid,
		"commentId": cid,
	})
	w = httptest.NewRecorder()

	service.DeleteComment(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	////////////////////////////////////////////////////
	// vote up and down. A few positive tests
	voteUp := &posts.Vote{User: uid, Vote: 1}
	voteDown := &posts.Vote{User: uid, Vote: -1}
	postsRepo.EXPECT().Vote(pid, *voteUp).Return(resultPost, nil)
	postsRepo.EXPECT().Vote(pid, *voteDown).Return(resultPost, nil)
	postsRepo.EXPECT().Unvote(pid, uid).Return(resultPost, nil)

	req = httptest.NewRequest("POST", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.Upvote(w, req)
	service.Downvote(w, req)
	service.Unvote(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	bts, _ = json.Marshal(resultPost)
	img, _ = bts, 0
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// vote. repo err
	postsRepo.EXPECT().Vote(pid, *voteUp).Return(nil, errors.New("DB Error"))
	postsRepo.EXPECT().Vote(pid, *voteDown).Return(nil, errors.New("DB Error"))
	postsRepo.EXPECT().Unvote(pid, uid).Return(nil, errors.New("DB Error"))

	req = httptest.NewRequest("POST", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, &session.Session{
		ID:      sid,
		UserID:  uid,
		Expires: time.Now().Add(time.Hour),
	}))
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.Upvote(w, req)
	service.Downvote(w, req)
	service.Unvote(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`DB Error`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}

	// vote. no session err
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": pid,
	})
	w = httptest.NewRecorder()

	service.Upvote(w, req)
	service.Downvote(w, req)
	service.Unvote(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = []byte(`InternalServerError`)
	if !bytes.Contains(body, img) {
		t.Errorf("Invalid.\n%s\n%s", string(body), img)
		return
	}
}
