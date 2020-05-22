package handlers

import (
	"encoding/json"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/users"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/utils"
	"net/http"
)

// UsersHandler struct contains necessary attributes to handle Users
type UsersHandler struct {
	// Tmpl     *template.Template
	// Logger   *zap.SugaredLogger
	UsersRepo *users.UsersRepo
	// Sessions  *session.SessionsManager
}

type retObj struct {
	Token string `json:"token"`
}

// Login takes care about loggin a User in
func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {

	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString("couldn't parse user's information")
		http.Error(w, jsonMessage, http.StatusBadRequest)
		return
	}

	// log.Println("Username is " + user.Username)

	u, err := h.UsersRepo.Authorize(user.Username, user.Password)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString(err.Error())
		http.Error(w, jsonMessage, http.StatusUnauthorized)
		return
	}

	result, _ := json.Marshal(retObj{Token: u.Token})
	w.Write(result)
	// http.Redirect(w, r, "/", 302)
}

// Register creates a new User in the repository
func (h *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonMessage := utils.GetJSONMessageAsString("couldn't parse user's information")
		http.Error(w, jsonMessage, http.StatusBadRequest)
		return
	}

	_, err = h.UsersRepo.Register(&user)
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
