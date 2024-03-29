## gctl report

Uses Go tooling to generate reports on a go module

### Synopsis

Uses Go tooling to generate reports on a go module.
It uses the following commands: gofmt, go vet, gocyclo, ineffassign, and misspell

You must have these tools installed before running:
  - gofmt: installed with go
  - go vet: installed with go
  - gocyclo: go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
  - ineffassign: go install github.com/gordonklaus/ineffassign@latest
  - misspell: go install github.com/client9/misspell/cmd/misspell@latest


```
gctl report [flags]
```

### Options

```
  -h, --help   help for report
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gctl.yaml)
```

### SEE ALSO

* [gctl](gctl.md)	 - gctl (Galactus Controller) is the built-in CLI for managing everything in Galactus

###### Auto generated by spf13/cobra on 12-Nov-2023
