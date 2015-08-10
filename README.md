## graphc
### A utility for managing docker storage backends

####*Installation:*
`go get github.com/willmtemple/graphc`

This will build the binary and place it in `${GOPATH}/bin`. Alternatively, you
might check-out the sources with git and manually build the program using `go
build`.

####*Invocation:*
`graphc --help`

The program has subcommands which are analogues of the public `Driver` API in
Docker's `graphdriver`.

This program is strictly a work-in-progress and can (and probably will) do
everything up to and including eat your laundry.
