// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// List all functions potentially needed by program

var stdout io.Writer = os.Stdout
var standardPackages = make(map[string]struct{})
var (
	dirFlag = flag.String("dir", "",
		`Packages path`)
	nostdFlag = flag.Bool("nostd", true,
		`Don't output standard package's functions`)
)

func init() {
	initStandardPackages()
}

func main() {
	flag.Parse()
	if err := listFuncs(*dirFlag, *nostdFlag, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "list-funcs: %s\n", err)
		os.Exit(1)
	}
}

func listFuncs(dir string, nostd bool, args []string) error {
	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax,
		Dir:  dir,
	}
	initial, err := packages.Load(cfg, args...)
	if err != nil {
		return err
	}
	if packages.PrintErrors(initial) > 0 {
		return fmt.Errorf("packages contain errors")
	}

	// Create and build SSA-form program representation.
	prog, _ := ssautil.AllPackages(initial, 0)
	prog.Build()

	funcs := ssautil.AllFunctions(prog)
	for fn := range funcs {
		pkgPath := getPackagePath((fn))
		if nostd && isStandardPackage(pkgPath) {
			continue
		}
		fmt.Fprintln(stdout, fn.String())
	}

	return nil
}

func initStandardPackages() {
	pkgs, err := packages.Load(nil, "std")
	if err != nil {
		panic(err)
	}

	for _, p := range pkgs {
		standardPackages[p.PkgPath] = struct{}{}
	}
}

func getPackagePath(fn *ssa.Function) string {
	var path string
	if fn.Pkg != nil {
		path = fn.Pkg.Pkg.Path()
	} else { // for shared funcs
		name := strings.TrimPrefix(fn.String(), "(")
		name = strings.TrimPrefix(name, "*")
		name = strings.TrimPrefix(name, "struct{")
		path = strings.Split(name, ".")[0]
	}
	return path
}

func isStandardPackage(pkg string) bool {
	_, ok := standardPackages[pkg]
	isGoTools := strings.HasPrefix(pkg, "golang") // for golang.org/x/tools
	return ok || isGoTools
}
