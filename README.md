# gocollect

Gocollect gives you the ability to collect the test coverage from a running application without the need of stopping it.

## Why?

Sometimes is not trivial to collect the coverage data after the run of your integration tests. For example if your app is running in a Pod on Kubernetes then stopping it will destroy the Pod and the collected coverage files. You can mount a Volume to persist the data, but this complicates the setup.

With `gocollect` you will be able to automatically add a `/collect` endpoint to your HTTP server to the application to emit the coverage data to your `GOCOVERDIR` directory.

The only caveat is that your app will need to be built with the `-covermode=atomic` flag (and obviously with the `-cover` as well). But since this will be a testing binary it shouldn't be a problem.

## Usage

For simple cases you can just import the package with a blank identifier:

```go
import _ "github.com/enrichman/gocollect"
```

When the `GOCOVERDIR` is not set this is a no-op, so it doesn't add any overhead. The handler will be added to the `/collect` path, but you can customize this with the `GOCOLLECT_PATH` environment variable.

If you are using a new router you will need to start from the `http.DefaultServeMux` 

```go
mux := http.DefaultServeMux
```

or you can manually add the handler to your router

```go
mux := http.NewServeMux()
mux.Handle("/collect", gocollect.Collect("mydir"))
```

You can check the [examples](examples/).
