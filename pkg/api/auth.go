package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ysv/pick/pkg/datastore"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *login) Sanitize() {
	l.Email = strings.ToLower(strings.TrimSpace(l.Email))
}

// GET /api/session
func (api *API) GetSession(w http.ResponseWriter, r *http.Request) error {
	userCount, err := api.database.CountUsers()
	if err != nil {
		return err
	}

	// if 0 users in database, dashboard is public
	if userCount == 0 {
		return respond(w, http.StatusOK, envelope{Data: true})
	}

	// if existing session, assume logged-in
	session, _ := api.sessions.Get(r, "auth")
	if !session.IsNew {
		return respond(w, http.StatusOK, envelope{Data: true})
	}

	// otherwise: not logged-in yet
	return respond(w, http.StatusOK, envelope{Data: false})
}

// URL: POST /api/session
func (api *API) CreateSession(w http.ResponseWriter, r *http.Request) error {
	// check login creds
	var l login
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		return err
	}
	l.Sanitize()

	// find user with given email
	u, err := api.database.GetUserByEmail(l.Email)
	if err != nil && err != datastore.ErrNoResults {
		return err
	}

	// compare pwd
	if err == datastore.ErrNoResults || u.ComparePassword(l.Password) != nil {
		return respond(w, http.StatusUnauthorized, envelope{Error: "invalid_credentials"})
	}

	// ignore error here as we want a (new) session regardless
	session, _ := api.sessions.Get(r, "auth")
	session.Values["user_id"] = u.ID
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return respond(w, http.StatusOK, envelope{Data: true})
}

// URL: DELETE /api/session
func (api *API) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := api.sessions.Get(r, "auth")
	if !session.IsNew {
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			return err
		}
	}

	return respond(w, http.StatusOK, envelope{Data: true})
}
