/*
Copyright (c) 2018, Sylabs, Inc. All rights reserved.
This software is licensed under a 3-clause BSD license.  Please
consult LICENSE file distributed with the sources of this project regarding
your rights to use or distribute this software.
*/

package main

import (
	"github.com/singularityware/singularity/internal/pkg/cli"
	"github.com/singularityware/singularity/pkg/build"
)

func main() {
	cli.Execute()
	b := build.NewLocalBuilder([]byte{})
	b.Build()
}