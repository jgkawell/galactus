[![Go Report Card](https://goreportcard.com/badge/github.com/jgkawell/galactus/pkg/chassis)](https://goreportcard.com/report/github.com/jgkawell/galactus/pkg/chassis)

# chassis

TODO: Write a project description

## Testing

Be sure to run tests as:

```shell
go test -gcflags=-l -test.v ./...

# on ARM-based machines (e.g. Apple Silicon)
GOARCH=amd64 go test -gcflags=-l -test.v ./...
```

## Mocks

Generate mocks using `mockery`:

```shell
mockery --all --inpackage --case underscore
```
