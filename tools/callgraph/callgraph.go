package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/token"
	"io"
	"os"
	"text/template"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa/ssautil"
)

var stdout io.Writer = os.Stdout
var (
	modFlag = flag.String("mod", "",
		"Use mod like build system.")
	focusFlag = flag.String("focus", "",
		"Focus on these packages, separated by comma, only output call graph in these packages.")
)

func main() {
	flag.Parse()
	initCheckCond(*focusFlag)
	if err := doCallgraph("", *modFlag, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "callgraph: %s\n", err)
		os.Exit(1)
	}
}

// Generate callgraph with pointer.Analyze
func doCallgraph(dir, mod string, args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, Usage)
		return nil
	}

	var buildFlags []string
	if len(mod) > 0 {
		modArg := fmt.Sprintf("-mod=%s", mod)
		buildFlags = []string{modArg}
	}
	cfg := &packages.Config{
		Mode:       packages.LoadAllSyntax,
		Dir:        dir,
		BuildFlags: buildFlags,
	}
	initial, err := packages.Load(cfg, args...)
	if err != nil {
		return err
	}
	if packages.PrintErrors(initial) > 0 {
		return fmt.Errorf("packages contain errors")
	}

	// Create and build SSA-form program representation.
	prog, pkgs := ssautil.AllPackages(initial, 0)
	prog.Build()

	// -- call graph construction ------------------------------------------

	var cg *callgraph.Graph
	mains := ssautil.MainPackages(pkgs)
	config := &pointer.Config{
		Mains:          mains,
		BuildCallGraph: true,
	}
	ptares, err := pointer.Analyze(config)
	if err != nil {
		return err // internal error in pointer analysis
	}
	cg = ptares.CallGraph
	cg.DeleteSyntheticNodes()

	err = output(cg, prog.Fset)
	if err != nil {
		return err
	}

	return nil
}

// Generate graphviz dot format output
func output(cg *callgraph.Graph, fset *token.FileSet) error {
	before := "digraph callgraph {\n"
	after := "}\n"
	format := `  {{printf "%q" .Caller}} -> {{printf "%q" .Callee}}`

	tmpl, err := template.New("-format").Parse(format)
	if err != nil {
		return fmt.Errorf("invalid -format template: %v", err)
	}

	// Allocate these once, outside the traversal.
	var buf bytes.Buffer
	data := Edge{fset: fset}

	fmt.Fprint(stdout, before)

	outputEdge := func(edge *callgraph.Edge) error {

		caller := edge.Caller
		callee := edge.Callee

		if isInitFunc(caller.Func.Name()) && isInitFunc(callee.Func.Name()) {
			return nil
		}
		if caller.Func.Pkg != nil && isStandardPkg(caller.Func.Pkg.Pkg.Path()) {
			return nil
		}
		if caller.Func.Pkg != nil && !isFocus(caller.Func.Pkg.Pkg.Path()) {
			return nil
		}

		data.position.Offset = -1
		data.edge = edge
		data.Caller = caller.Func
		data.Callee = callee.Func

		buf.Reset()
		if err := tmpl.Execute(&buf, &data); err != nil {
			return err
		}
		stdout.Write(buf.Bytes())
		if len := buf.Len(); len == 0 || buf.Bytes()[len-1] != '\n' {
			fmt.Fprintln(stdout)
		}
		return nil
	}

	if err := callgraph.GraphVisitEdges(cg, outputEdge); err != nil {
		return err
	}
	fmt.Fprint(stdout, after)
	return nil
}
