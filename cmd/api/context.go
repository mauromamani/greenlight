package main

import (
	"context"
	"net/http"

	"github.com/mauromamani/greenlight/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")

// contextSetUser:
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// contextGetUser:
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
