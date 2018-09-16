package handlers

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// Standard http handler
type handlerFunc func(w http.ResponseWriter, r *http.Request)

// Our context.Context http handler
type contextHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

// Higher order function for changing a HandlerFunc to a ContextHandlerFunc,
// usually creating the context.Context along the way.
type ContextHandlerToHandlerHOF func(f contextHandlerFunc) handlerFunc

// Creates a new Context and uses it when calling f
func ContextHandlerToHTTPHandler(f contextHandlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		f(ctx, w, r)
	}
}

func ConvertTestingContextHandlerToHTTPHandler(ctx context.Context) ContextHandlerToHandlerHOF {
	return func(f contextHandlerFunc) handlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(ctx, w, r)
		}
	}
}

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
