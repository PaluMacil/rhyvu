package ry

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"io/ioutil"

	"github.com/satori/go.uuid"
)

// Session represents a user session
type Session struct {
	Token     uuid.UUID `storm:"id"`
	User      User      `storm:"index"`
	Created   time.Time
	HeartBeat time.Time
}

// TokenRequest holds the incoming request for a token (session key)
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TokenHandler takes in unencrypted credentials (email and password) via form values,
// hashes the password with bcrypt, compares to the stored hash, and returns Unauthorized
// or a token (uuid version 4) representing the session.
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("could not read login request json body:", err)
		http.Error(w, "could not read login request json body", http.StatusInternalServerError)
		return
	}
	var req TokenRequest
	err = json.Unmarshal(body, &req)

	var user User
	err = Db.One("Email", req.Email, &user)
	if err != nil {
		log.Println("could not load user to build session:", err, "for email", req.Email)
		http.Error(w, "incorrect email or password", http.StatusBadRequest)
		return
	}
	if CheckPasswordHash(req.Password, user.Password) {
		session := Session{
			Token:     uuid.NewV4(),
			User:      user,
			Created:   time.Now(),
			HeartBeat: time.Now(),
		}
		err = Db.Save(&session)
		if err != nil {
			log.Println("could not save session:", err)
			http.Error(w, "could not save session", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(session)
	} else {
		log.Println("incorrect password:", req.Password, "for user:", user.Email)
		http.Error(w, "incorrect email or password", http.StatusBadRequest)
		return
	}
}

// SessionFor returns the session for the current request
func SessionFor(r *http.Request) (Session, error) {
	token := r.Header.Get("X-DWN-TOKEN")
	var session Session
	if err := Db.One("Token", token, &session); err != nil {
		return session, err
	}
	return session, nil
}

// LogoutHandler handles requests to log out. If the session exists, it is deleted.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-DWN-TOKEN")
	var session Session
	if err := Db.One("Token", token, &session); err != nil {
		http.Error(w, "Could not find session.", http.StatusBadRequest)
		return
	}
	if err := Db.DeleteStruct(&session); err != nil {
		log.Println("Could not delete session:", err)
		http.Error(w, "Could not delete session.", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
