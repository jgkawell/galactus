# clientfactory

This package provides a factory for creating RPC clients.

## Generate Mocks

To generate mocks run the Makefile target replacing the interface with whichever interface you want to mock:

```sh
make mockclient AGGREGATE=soccer VERSION=1
```
Default VERSION is 1