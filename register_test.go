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
	ctx, inst := getContext()
	defer inst.Close()

	c.Convey("When visit the signup page", t, func() {
		r := createHTTPRouter(createContextHandlerToHTTPHandler(ctx))
		record := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/signup", nil) // URL-encoded payload

		c.So(err, c.ShouldBeNil)

		c.Convey("The next page body should contain \"Please enter your email address\"", func() {
			r.ServeHTTP(record, req)
			c.So(record.Code, c.ShouldEqual, http.StatusOK)
		})
	})
}
