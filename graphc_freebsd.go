// +build freebsd

package main

import (
	_ "github.com/willmtemple/graphc/graphdriver/zfs"
)

var (
	defaultHome = "/var/lib/docker"
)
