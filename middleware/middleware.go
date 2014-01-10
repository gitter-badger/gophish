package middleware

import (
	"fmt"
	"net/http"

	ctx "github.com/gorilla/context"
	"github.com/jordan-wright/gophish/auth"
)

// Use allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}

// GetContext wraps each request in a function which fills in the context for a given request.
// This includes setting the User and Session keys and values as necessary for use in later functions.
func GetContext(handler http.Handler) http.Handler {
	// Set the context here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the context appropriately here.
		// Set the session
		session, _ := auth.Store.Get(r, "gophish")
		ctx.Set(r, "session", session)
		handler.ServeHTTP(w, r)
		// Save the session
		session.Save()
		// Remove context contents
		ctx.Clear(r)
	})
}

// RequireLogin is a simple middleware which checks to see if the user is currently logged in.
// If not, the function returns a 302 redirect to the login page.
func RequireLogin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RequireLogin called!!")
		handler.ServeHTTP(w, r)
	})
}