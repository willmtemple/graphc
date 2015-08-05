// +build windows

package main

import (
	"os"

	_ "github.com/willmtemple/graphc/graphdriver/windows"
)

var (
	defaultHome = os.Getenv("programdata") + string(os.PathSeparator) + "docker"
)
