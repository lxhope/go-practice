// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

// List all packages potentially needed by program

var stdout io.Writer = os.Stdout
var (
	dirFlag = flag.String("dir", "",
		`Packages path`)
)

func main() {
	flag.Parse()
	if err := listPkgs(*dirFlag, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "list-pkgs: %s\n", err)
		os.Exit(1)
	}
}

func listPkgs(dir string, args []string) error {
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

	for _, pkg := range prog.AllPackages() {
		fmt.Fprintln(stdout, pkg.Pkg.Name(), pkg.Pkg.Path())
	}

	return nil
}
