# aecontext-handlers

aecontext-handlers are a bunch of AppEngine context handlers that wrap the
standard `net/http` HandlerFunc, enabling you to easily switch out how the
AppEngine Context is created, depending on what environment (production / test)
the code is running in.

## Example usage

Note that this is an AppEngine main application, hence there is no `func main() {}`.

### `main.go`

```go
package main

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"

	"github.com/tintinnabulate/aecontext-handlers/handlers"
)

// createHTTPRouter : creates our URL router. As an argument it takes in a
// ToHandlerHOF, whose job is to make the conversion to a standard HandlerFunc()
// format. This means we can change how the AppEngine Context gets created, by
// passing in different ToHandlerHOF implementations. In this case, we use the one
// we want for production (see `init()` function)
func createHTTPRouter(f handlers.ToHandlerHOF) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/foo", f(fooHandler)).Methods("GET")
	return appRouter
}

// fooHandler : handle /foo
func fooHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
}

func init() {
    // create our URL router, passing in a HOF that converts to a standard HandlerFunc()
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
	"github.com/tintinnabulate/aecontext-handlers/handlers"
)

// TestFooHandler : does just that
func TestFooHandler(t *testing.T) {
    // Get our testing context and aetest instance
	ctx, inst := handlers.GetTestingContext()
    // Remember to close it at the end of every test!
	defer inst.Close()

	c.Convey("When user visits the foo page", t, func() {
        // create our router, using this testing context
		r := createHTTPRouter(handlers.ToHTTPHandlerConverter(ctx))
		record := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/foo", nil)
		c.So(err, c.ShouldBeNil)

		c.Convey("The status should equal http.StatusOK", func() {
            // use the router to serve the request, recording the response
			r.ServeHTTP(record, req)
            // Assert that the return code matches what we expect
			c.So(record.Code, c.ShouldEqual, http.StatusOK)
		})
	})
}
```

## Used by

* <https://github.com/tintinnabulate/registration-webapp> (See `register.go` and `register_test.go`)
* <https://github.com/tintinnabulate/vmail> (See `signup.go` and `signup_test.go`)

## Inspired by

This idea was taken, almost completely unmodified, from the excellent Compound Theory blog:

<https://www.compoundtheory.com/testing-go-http-handlers-in-google-app-engine-with-mux-and-higher-order-functions/>

All credit goes to that :-)
