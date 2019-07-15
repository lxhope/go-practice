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
	"golang.org/x/tools/go/ssa/ssautil"
)

// List all packages potentially needed by program

var stdout io.Writer = os.Stdout
var standardPackages = make(map[string]struct{})
var (
	dirFlag = flag.String("dir", "",
		`Packages path`)
	nostdFlag = flag.Bool("nostd", true,
		`Don't output standard package's functions`)
	modFlag = flag.String("mod", "",
		"Use mod like build system.")
)

func init() {
	initStandardPackages()
}

func main() {
	flag.Parse()
	if err := listPkgs(*dirFlag, *nostdFlag, *modFlag, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "list-pkgs: %s\n", err)
		os.Exit(1)
	}
}

func listPkgs(dir string, nostd bool, mod string, args []string) error {
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
	prog, _ := ssautil.AllPackages(initial, 0)
	prog.Build()

	for _, pkg := range prog.AllPackages() {
		if nostd && isStandardPackage(pkg.Pkg.Path()) {
			continue
		}
		fmt.Fprintln(stdout, pkg.Pkg.Path())
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

func isStandardPackage(pkg string) bool {
	_, ok := standardPackages[pkg]
	isGoTools := strings.HasPrefix(pkg, "golang") // for golang.org/x/tools
	return ok || isGoTools
}
