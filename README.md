# http-mock [![CodeFactor](https://www.codefactor.io/repository/github/claudemuller/http-mock/badge)](https://www.codefactor.io/repository/github/claudemuller/http-mock)
A mock http server that returns dummy responses.

## Adding Routes
Nest a list of `Route`s into a list of `ParentRoute`s. The `Path` variable of a `ParentRoute` instance creates a `SubRouter` on the `mux` `Router`.

### Route Values in `config.go`
#### `PreRouter struct`
- `Path` - the path of the parent route segment
- `SubRoutes` - a list of sub-`Route`s on the above parent route segment

#### `Router struct`
- `path` - the path of the endpoint to mock
- `param` - the parameter to capture, if any
- `method` - the HTTP method to handle
- `contentType` - the content type to return
- `responseStatus` -  the HTTP response code
- `response` - the response to send (atm only JSON)

## Installing Dependencies
`go install`

## Building
`go build`

## Running
`go run`

## License
http-mock is [MIT licensed](https://github.com/claudemuller/http-mock/new/master?filename=LICENSE.md).
