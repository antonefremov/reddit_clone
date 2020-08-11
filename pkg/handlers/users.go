package handlers

import (
	"encoding/json"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/ljwt"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/session"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/user"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/utils"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// UsersRepoInterface describes the set of functions available for the Users Repo
type UsersRepoInterface interface {
	GetByID(string) (*user.User, error)
	GetByUserName(string) (*user.User, error)
	Register(*user.User) (*user.User, error)
	Authorize(string, string) (*user.User, error)
}

// SessionsManagerInterface partially defines interface of the SM
type SessionsManagerInterface interface {
	Create(string) (*session.Session, error)
}

// UsersHandler struct contains necessary attributes to handle Users
type UsersHandler struct {
	// Tmpl     *template.Template
	Logger    *zap.SugaredLogger
	UsersRepo UsersRepoInterface
	Sessions  SessionsManagerInterface
}

type loginForm struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}

// RetObj is the object expected by the front end app and being returned
type RetObj struct {
	Token string `json:"token"`
}

// Login takes care about loggin a User in
func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	lf := &loginForm{}
	err = json.Unmarshal(request, lf)
	if err != nil {
		h.Logger.Errorf(`BadRequest. %s`, err.Error())
		http.Error(w, `BadRequest. Can't parse request body'`, http.StatusBadRequest)
		return
	}

	uAuth, err := h.UsersRepo.Authorize(lf.Login, lf.Password)
	if err == user.ErrNoUser {
		h.Logger.Errorf(`BadRequest. %s`, err.Error())
		jsonMessage(w, http.StatusBadRequest, "user not found")
		return
	}
	if err == user.ErrBadPass {
		h.Logger.Errorf(`BadRequest. %s`, err.Error())
		jsonMessage(w, http.StatusBadRequest, "invalid password")
		return
	}
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	// Session
	sess, err := h.Sessions.Create(uAuth.ID)
	if err != nil {
		h.Logger.Errorf(`InternalServerError. %s`, err.Error())
		http.Error(w, `InternalServerError`, http.StatusInternalServerError)
		return
	}

	// generate a JWT token
	tokenString, err := ljwt.IssueNewToken(uAuth.ID, uAuth.Username, sess.ID)

	result, _ := json.Marshal(RetObj{Token: tokenString})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	// http.Redirect(w, r, "/", 302)
}

// Register creates a new User in the repository
func (h *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString("couldn't parse user's information")
		http.Error(w, jsonMessage, http.StatusBadRequest)
		return
	}

	_, err = h.UsersRepo.Register(&u)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/api/login", 307)
}

// Logout destroys User's credentials
func (h *UsersHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// h.Sessions.DestroyCurrent(w, r)
	http.Redirect(w, r, "/", 302)
}

func jsonMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"message":"` + message + `"}`))
}
