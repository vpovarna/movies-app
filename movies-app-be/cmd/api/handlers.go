package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.GetAllMovies()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	// read a json payload (username and password)
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.ReadJson(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// validate user against DB
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FistName,
		LastName:  user.LastName,
	}

	// generate token
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)
}

func (app *application) RefreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			//parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})

			if err != nil {
				app.errorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the userId from the token claims
			userId, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJson(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserById(userId)
			if err != nil {
				app.errorJson(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:        userId,
				FirstName: user.FistName,
				LastName:  user.LastName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJson(w, errors.New("error generating tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
}
