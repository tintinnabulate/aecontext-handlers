# aecontext-handlers
aecontext-handlers are a bunch of AppEngine context handlers that wrap the standard `net/http` HandlerFunc, enabling you to fully test each HTTP Handler by passing in the correct context for live/test mode.

* In Live mode: AppEngine context is passed into HandlerFunc

* In Test mode: standard context.Context is passed into HandlerFunc

## Example usage

foo.go:

```go
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
func createHTTPRouter(f handlers.ToHandlerHOF) *mux.Router {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/signup", f(getSignupHandler)).Methods("GET")
	return appRouter
}

// getSignupHandler : show the signup form (SignupURL)
func getSignupHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
}

func init() {
	router := createHTTPRouter(handlers.ToHTTPHandlerConverter)
	http.Handle("/", csrfProtectedRouter)
}
```

foo_test.go:

```go
package main

import (
    "fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

// TestGetSignupPage does just that
func TestGetSignupPage(t *testing.T) {
	ctx, inst := handlers.GetTestingContext()
	defer inst.Close()

	c.Convey("When visit the signup page", t, func() {
		r := createHTTPRouter(handlers.ToHTTPHandler(ctx))
		record := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/signup", nil) // URL-encoded payload

		c.So(err, c.ShouldBeNil)

		c.Convey("The next page body should contain \"Please enter your email address\"", func() {
			r.ServeHTTP(record, req)
			c.So(record.Code, c.ShouldEqual, http.StatusOK)
		})
	})
}
```

