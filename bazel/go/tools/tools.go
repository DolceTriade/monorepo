//go:build tools

package main

import (
	_ "github.com/kisielk/errcheck/errcheck"
	_ "honnef.co/go/tools/analysis/lint"
)
