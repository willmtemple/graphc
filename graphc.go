package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/willmtemple/graphc/graphdriver"

	_ "github.com/willmtemple/graphc/graphdriver/aufs"
	_ "github.com/willmtemple/graphc/graphdriver/devmapper"
	_ "github.com/willmtemple/graphc/graphdriver/vfs"
	_ "github.com/willmtemple/graphc/graphdriver/overlay"
	_ "github.com/willmtemple/graphc/graphdriver/btrfs"
)

func initDriver(c *cli.Context) graphdriver.Driver {
	graphdriver.DefaultDriver = c.GlobalString("driver")
	homedir := c.GlobalString("home")
	drv, err := graphdriver.New(homedir, []string{})
	if err != nil {
		fmt.Printf("Failed to instantiate graphdriver: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("[DEBUG] Using driver %s.\n%g\nHome directory: %s\n", drv.String(), drv.Status(), homedir)
	return drv
}

func create(c *cli.Context) {
	driver := initDriver(c)
	id := c.Args().First()
	if err := driver.Create(id, c.String("parent")); err != nil {
		fmt.Printf("Failed to create %s: %s\n", id, err)
		driver.Cleanup()
		os.Exit(1)
	}
}

func get(c *cli.Context) {
	driver := initDriver(c)
	id := c.Args().First()
	loc, err := driver.Get(id, c.GlobalString("context"))
	if err != nil {
		fmt.Printf("Failed to Get %s: %s\n", id, err)
		driver.Cleanup()
		os.Exit(1)
	}
	fmt.Printf("%s is available at %s\n", id, loc)
}

func put(c *cli.Context) {
	driver := initDriver(c)
	id := c.Args().First()
	if err := driver.Put(id); err != nil {
		fmt.Printf("Failed to Put %s: %s\n", id, err)
		driver.Cleanup()
		os.Exit(1)
	}
}

func main() {

	graphc := cli.NewApp()
	graphc.Name = "graphc"
	graphc.Usage = "manage graphc storage"
	graphc.Flags = []cli.Flag {
		cli.StringFlag{
			Name:	"home",
			Value:	"/var/lib/docker/",
			Usage:	"home directory for graphdriver storage operations",
			EnvVar:	"GRAPHDRIVER_HOME",
		},
		cli.StringFlag{
			Name:	"driver, s",
			Value:	"",
			Usage:	"storage backend to use",
			EnvVar:	"GRAPHDRIVER_BACKEND",
		},
		cli.StringFlag{
			Name:	"context, c",
			Value:	"",
			Usage:	"optional mountlabel (SELinux context)",
		},

	}
	graphc.EnableBashCompletion = true
	graphc.Commands = []cli.Command{
		{
			Name:		"create",
			Aliases:	[]string{"c"},
			Usage:		"create a new storage for id",
			Flags:		[]cli.Flag {
				cli.StringFlag{
					Name:	"parent, p",
					Value:	"",
					Usage:	"an id of which the new image will initially be a copy",
				},
			},
			Action:		create,
		},
		{
			Name:		"get",
			Aliases:	[]string{"g"},
			Usage:		"mount an image to the filesystem",
			Action:		get,
		},
		{
			Name:		"put",
			Aliases:	[]string{"p"},
			Usage:		"unmount an image from the filesystem",
			Action:		put,
		},
	}

	graphc.Run(os.Args)

}
