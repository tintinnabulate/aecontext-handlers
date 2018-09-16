# aecontext-handlers
aecontext-handlers are a bunch of AppEngine context handlers that wrap the standard `net/http` HandlerFunc, enabling you to fully test each HTTP Handler by passing in the correct context for live/test mode.

* In Live mode: AppEngine context is passed into HandlerFunc

* In Test mode: standard context.Context is passed into HandlerFunc

## Example usage

Foo
