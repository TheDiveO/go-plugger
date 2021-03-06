# plugger

[![GoDoc](https://godoc.org/github.com/TheDiveO/go-plugger?status.svg)](http://godoc.org/github.com/TheDiveO/go-plugger)
![GitHub](https://img.shields.io/github/license/thediveo/go-asciitree)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/go-plugger)](https://goreportcard.com/report/github.com/thediveo/go-plugger)

`plugger` is a simplistic plugin manager that works with both statically
linked as well as dynamically linked Go plugins. It supports multiple plugin
groups, as well as controlled plugin order within a group.

## Examples

See the `examples` directory for how to use plugger in a Go application for
organizing and using static plugins (plugins that have been statically linked
into your application). `examples/staticplugins/main.go` uses plugger to
get all plugins in the "plugin" group and then calls some method on them.

```go
import (
    // import your plugins
    _ "github.com/thediveo/go-plugger/examples/staticplugins/plugins/foo"
    _ "..."
)

plugs := plugger.New("plugins")
for _, doit := range plugs.Func("DoIt") {
    fmt.Println(doit.(func() string)())
}
```

The plugins get statically linked in by importing them, such as `plugins/foo`.
While at first this might seem to be much overhead, the more plugins you have
in your application, and the more groups you need them to organize into, the
more you'll benefit from the `go-plugger` package: you only need to import
the plugin packages, and plugger will do the rest.

For a more elaborate "example", please also look at `internal/staticplugin/`
and `internal/dynamicplugin/` (these are the built-in test cases).

## Run Unit Tests

- VisualStudio Code: please first build the workspace, before running
  tests.

- from CLI: simply run the `./testall.sh` script; it will build the shared
  library for the test shared library plugin.

## Copyright and License

`plugger` is Copyright 2019 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
