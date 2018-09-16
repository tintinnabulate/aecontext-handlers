package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go"
	stripeClient "github.com/stripe/stripe-go/client"

	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"
)

// createHTTPRouter : create a HTTP router where each handler is wrapped by a given context
func createHTTPRouter(f contextHandlerToHandlerHOF) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/signup", f(getSignupHandler)).Methods("GET")
	return appRouter
}

// getSignupHandler : show the signup form (SignupURL)
func getSignupHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
}

func init() {
	router := createHTTPRouter(contextHandlerToHTTPHandler)
	http.Handle("/", csrfProtectedRouter)
}
