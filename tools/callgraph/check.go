package main

import (
	"strings"

	"golang.org/x/tools/go/packages"
)

var standardPackages = make(map[string]struct{})
var focusPackages = []string{}

// initial check condition
func initCheckCond(focusPkgs string) {
	initStandardPkgs()
	if len(focusPkgs) > 0 {
		initFocus(focusPkgs)
	}
}

// initial standardPackages
func initStandardPkgs() {
	pkgs, err := packages.Load(nil, "std")
	if err != nil {
		panic(err)
	}

	for _, p := range pkgs {
		standardPackages[p.PkgPath] = struct{}{}
	}
}

// initial focusPackages
func initFocus(focusPkgs string) {
	if len(focusPackages) == 0 && len(focusPkgs) > 0 {
		pkgs := strings.Split(focusPkgs, ",")
		for _, item := range pkgs {
			focusPackages = append(focusPackages, item)
		}
	}
}

func isStandardPkg(pkg string) bool {
	_, ok := standardPackages[pkg]
	isGoTools := strings.HasPrefix(pkg, "golang.org")
	return ok || isGoTools
}

func isInitFunc(name string) bool {
	return strings.HasSuffix(name, "init")
}

func isFocus(pkg string) bool {
	if len(focusPackages) == 0 { // no focus args,
		return true
	}
	for _, prefix := range focusPackages {
		if strings.HasPrefix(pkg, prefix) {
			return true
		}
	}
	return false
}
