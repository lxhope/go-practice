package main

const Usage = `callgraph: display the the call graph of a Go program.

Usage:

  callgraph [-mod=vendor] [-focus=...] package...

Flags:

-mod       Delivery the mod argument to build tools.

-focus     Specifies the package's prefix in which each call graph edge is displayed.

Examples:

  Show the call graph of the trivial web server application:

    callgraph -mod=vendor $GOROOT/src/net/http/triv.go
`
