package main

import (
	"strings"

	"golang.org/x/tools/go/packages"
)

var standardPackages = make(map[string]struct{})

func init() {
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
	isGoTools := strings.HasPrefix(pkg, "golang.org")
	return ok || isGoTools
}

func isInitFunc(name string) bool {
	return strings.HasSuffix(name, "init")
}
