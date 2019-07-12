package main

const Usage = `callgraph: display the the call graph of a Go program.

Usage:

  callgraph [-algo=static|cha|rta|pta] [-test] [-format=...] package...

Flags:

-algo      Specifies the call-graph construction algorithm, one of:

            static      static calls only (unsound)
            cha         Class Hierarchy Analysis
            rta         Rapid Type Analysis
            pta         inclusion-based Points-To Analysis

           The algorithms are ordered by increasing precision in their
           treatment of dynamic calls (and thus also computational cost).
           RTA and PTA require a whole program (main or test), and
           include only functions reachable from main.

-test      Include the package's tests in the analysis.

-format    Specifies the format in which each call graph edge is displayed.
           One of:

            digraph     output suitable for input to
                        golang.org/x/tools/cmd/digraph.
            graphviz    output in AT&T GraphViz (.dot) format.

           All other values are interpreted using text/template syntax.
           The default value is:

            {{.Caller}}\t--{{.Dynamic}}-{{.Line}}:{{.Column}}-->\t{{.Callee}}

           The structure passed to the template is (effectively):

                   type Edge struct {
                           Caller      *ssa.Function // calling function
                           Callee      *ssa.Function // called function

                           // Call site:
                           Filename    string // containing file
                           Offset      int    // offset within file of '('
                           Line        int    // line number
                           Column      int    // column number of call
                           Dynamic     string // "static" or "dynamic"
                           Description string // e.g. "static method call"
                   }

           Caller and Callee are *ssa.Function values, which print as
           "(*sync/atomic.Mutex).Lock", but other attributes may be
           derived from them, e.g. Caller.Pkg.Pkg.Path yields the
           import path of the enclosing package.  Consult the go/ssa
           API documentation for details.

Examples:

  Show the call graph of the trivial web server application:

    callgraph -format digraph $GOROOT/src/net/http/triv.go

  Same, but show only the packages of each function:

    callgraph -format '{{.Caller.Pkg.Pkg.Path}} -> {{.Callee.Pkg.Pkg.Path}}' \
      $GOROOT/src/net/http/triv.go | sort | uniq

  Show functions that make dynamic calls into the 'fmt' test package,
  using the pointer analysis algorithm:

    callgraph -format='{{.Caller}} -{{.Dynamic}}-> {{.Callee}}' -test -algo=pta fmt |
      sed -ne 's/-dynamic-/--/p' |
      sed -ne 's/-->.*fmt_test.*$//p' | sort | uniq

  Show all functions directly called by the callgraph tool's main function:

    callgraph -format=digraph golang.org/x/tools/cmd/callgraph |
      digraph succs golang.org/x/tools/cmd/callgraph.main
`
