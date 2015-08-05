// +build linux

package main

import (
	_ "github.com/willmtemple/graphc/graphdriver/aufs"
	_ "github.com/willmtemple/graphc/graphdriver/btrfs"
	_ "github.com/willmtemple/graphc/graphdriver/devmapper"
	_ "github.com/willmtemple/graphc/graphdriver/overlay"
	_ "github.com/willmtemple/graphc/graphdriver/zfs"
)

var (
	defaultHome = "/var/lib/docker"
)
