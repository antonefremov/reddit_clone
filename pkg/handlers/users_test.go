package handlers

import (
	"bytes"
	"errors"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/session"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/user"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestHandlerRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sessRepo := NewMockSessionsManagerInterface(ctrl)
	userRepo := NewMockUsersRepoInterface(ctrl)
	service := UsersHandler{
		UsersRepo: userRepo,
		Logger:    zap.NewNop().Sugar(),
		Sessions:  sessRepo,
	}

	login := "login_test"
	uid := "id_test"
	pass := "password_test"

	registeringUser := &user.User{
		Username: login, Password: pass,
	}

	resultUser := &user.User{
		Username: login, ID: uid,
	}

	userRepo.EXPECT().Register(registeringUser).Return(resultUser, nil)

	/////////////////////////////////////////////////////////////////////////
	// Valid request -> 307 (Temporary Redirect)
	req := httptest.NewRequest("GET", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w := httptest.NewRecorder()

	service.Register(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	msg := `Temporary Redirect`
	if !bytes.Contains(body, []byte(msg)) {
		t.Errorf("Invalid returned result: %s", string(body))
		return
	}

	/////////////////////////////////////////////////////////////////////////
	// No body -> 'Bad Request'
	req = httptest.NewRequest("POST", "/", nil)
	w = httptest.NewRecorder()

	service.Register(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	msg = `{"message":"couldn't parse user's information"}`
	if !bytes.Contains(body, []byte(msg)) {
		t.Errorf("Invalid returned result: %s", string(body))
		return
	}

	// ///////////////////////////////////////
	// // register repo error already registered
	userRepo.EXPECT().Register(registeringUser).Return(nil, user.ErrUserExists)
	req = httptest.NewRequest("POST", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w = httptest.NewRecorder()

	service.Register(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	msg = `{"message":"User name already exists"}`
	if !bytes.Contains(body, []byte(msg)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}
}

func TestHandlerLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	sessRepo := NewMockSessionsManagerInterface(ctrl)
	userRepo := NewMockUsersRepoInterface(ctrl)
	service := UsersHandler{
		UsersRepo: userRepo,
		Logger:    zap.NewNop().Sugar(),
		Sessions:  sessRepo,
	}

	login := "login_test"
	uid := "id_test"
	sid := "sess_id_test"
	pass := "password_test"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVzZXJuYW1lIjoibG9naW5fdGVzdCIsImlkIjoiaWRfdGVzdCJ9LCJzZXNzaW9uSWQiOiJzZXNzX2lkX3Rlc3QifQ.IRgKxDeFVbsAh2VxZDNSiUzYsyJ2plkt6ziVVkpO6xE"

	resultUser := &user.User{
		Username: login, ID: uid,
	}
	resultSess := &session.Session{
		ID: sid, UserID: uid, Expires: time.Now(),
	}

	userRepo.EXPECT().Authorize(login, pass).Return(resultUser, nil)
	sessRepo.EXPECT().Create(uid).Return(resultSess, nil)

	req := httptest.NewRequest("POST", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w := httptest.NewRecorder()

	service.Login(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img := `{"token":"` + token + `"}`
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}

	/////////////////////////////////////////////////////////////////////////
	// No body
	req = httptest.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()

	service.Login(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = `BadRequest. Can't parse request body'`
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}

	///////////////////////////////////////
	// login repo error
	userRepo.EXPECT().Authorize(login, pass).Return(nil, user.ErrNoUser)
	req = httptest.NewRequest("GET", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w = httptest.NewRecorder()

	service.Login(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = `{"message":"user not found"}`
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}

	// ///////////////////////////////////////
	// login repo error
	userRepo.EXPECT().Authorize(login, pass).Return(nil, user.ErrBadPass)
	req = httptest.NewRequest("GET", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w = httptest.NewRecorder()

	service.Login(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = `{"message":"invalid password"}`
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}

	// ///////////////////////////////////////
	// // login repo error
	userRepo.EXPECT().Authorize(login, pass).Return(nil, errors.New(""))
	req = httptest.NewRequest("GET", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w = httptest.NewRecorder()

	service.Login(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = `InternalServerError`
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}

	// ///////////////////////////////////////
	// login sess repo error
	userRepo.EXPECT().Authorize(login, pass).Return(resultUser, nil)
	sessRepo.EXPECT().Create(uid).Return(nil, errors.New(""))
	req = httptest.NewRequest("GET", "/",
		strings.NewReader(`{"username":"`+login+`", "password": "`+pass+`"}`))
	w = httptest.NewRecorder()

	service.Login(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	img = `InternalServerError`
	if !bytes.Contains(body, []byte(img)) {
		t.Errorf("Invalid. %s", string(body))
		return
	}
}

func TestHandlerLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sessRepo := NewMockSessionsManagerInterface(ctrl)
	userRepo := NewMockUsersRepoInterface(ctrl)
	service := UsersHandler{
		UsersRepo: userRepo,
		Logger:    zap.NewNop().Sugar(),
		Sessions:  sessRepo,
	}

	/////////////////////////////////////////////////////////////////////////
	// Valid request -> 302 (Found)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	service.Logout(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	msg := `Found`
	if !bytes.Contains(body, []byte(msg)) {
		t.Errorf("Invalid returned result: %s", string(body))
		return
	}
}
