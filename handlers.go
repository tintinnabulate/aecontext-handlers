package handlers

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

// HandlerFunc : the standard http handler
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// ContextHandlerFunc : like HandlerFunc, plus the golang.org/x/net/context Context
type ContextHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

// ToHandlerHOF :
// Higher order function for changing a HandlerFunc to a ContextHandlerFunc,
// usually creating the context.Context along the way.
type ToHandlerHOF func(f ContextHandlerFunc) HandlerFunc

// ToHTTPHandler : returns a HandlerFunc which uses a new AppEngine context inside it
func ToHTTPHandler(f ContextHandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		f(ctx, w, r)
	}
}

// ToHTTPHandlerConverter : returns a higher order function that converts
// a testing context handler to a standard HTTP handler
func ToHTTPHandlerConverter(ctx context.Context) ToHandlerHOF {
	return func(f ContextHandlerFunc) HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(ctx, w, r)
		}
	}
}

// GetTestingContext : gets appengine context and appengine test instance
func GetTestingContext() (context.Context, aetest.Instance) {
	inst, _ := aetest.NewInstance(
		// TODO: pass these aetest.Options in as param to getTestingContext
		&aetest.Options{
			StronglyConsistentDatastore: true,
		})
	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		inst.Close()
	}
	ctx := appengine.NewContext(req)
	return ctx, inst
}
