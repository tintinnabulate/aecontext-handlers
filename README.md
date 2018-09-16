# aecontext-handlers
aecontext-handlers are a bunch of AppEngine context handlers that wrap the standard `net/http` HandlerFunc, enabling you to fully test each HTTP Handler by passing in the correct context for live/test mode.

## Example usage

This is an appengine main application (hence there is no `func main() {}` -- see `func init() {...}` for the definition of the http routes)

### `main.go`

```go
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

    handlers "github.com/tintinnabulate/aecontext-handlers"
)

// createHTTPRouter : create a HTTP router where each handler is wrapped by a given context
func createHTTPRouter(f handlers.ToHandlerHOF) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/foo", f(fooHandler)).Methods("GET")
	return appRouter
}

// fooHandler : handle /foo
func fooHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
}

func init() {
	router := createHTTPRouter(handlers.ToHTTPHandler)
	http.Handle("/", router)
}
```

### `main_test.go`

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/smartystreets/goconvey/convey"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// TestFooHandler : does just that
func TestFooHandler(t *testing.T) {
	ctx, inst := handlers.GetTestingContext()
	defer inst.Close()

	c.Convey("When user visits the foo page", t, func() {
		r := createHTTPRouter(handlers.ToHTTPHandlerConverter(ctx))
		record := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/foo", nil)
		c.So(err, c.ShouldBeNil)

		c.Convey("The status should equal http.StatusOK", func() {
			r.ServeHTTP(record, req)
			c.So(record.Code, c.ShouldEqual, http.StatusOK)
		})
	})
}
```

## Used by

* <https://github.com/tintinnabulate/registration-webapp> (See `register.go` and `register_test.go`)
