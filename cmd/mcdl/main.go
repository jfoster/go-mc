package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/jfoster/go-minecraft/version"
)

type context struct {
	client *version.Client
}

func (ctx *context) getlatest(snapshot bool) (string, error) {
	if snapshot {
		return ctx.client.GetLatestSnapshotID()
	}
	return ctx.client.GetLatestReleaseID()
}

type DownloadCmd struct {
	Snapshot bool `short:"s"`

	Version string `arg optional`
}

func (c *DownloadCmd) Run(ctx *context) (err error) {
	version, err := ctx.getlatest(c.Snapshot)

	if c.Version != "" {
		version = c.Version
	}

	fmt.Println(version)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, v, err := ctx.client.GetVersion(version)
	if err != nil {
		return err
	}

	return ctx.client.DownloadServerJar(v, pwd)
}

type VersionsCmd struct {
	Snapshot bool `short:"s"`
	Beta     bool `short:"b"`
	Alpha    bool `short:"a"`
}

func (c *VersionsCmd) Run(ctx *context) error {
	_, man, err := ctx.client.GetManifest()
	if err != nil {
		return err
	}

	var ids []version.ID

	for _, v := range man.Versions {
		if v.Type == version.ReleaseType ||
			(c.Snapshot && v.Type == version.SnapshotType) ||
			(c.Beta && v.Type == version.BetaType) ||
			(c.Alpha && v.Type == version.AlphaType) {
			ids = append(ids, v.ID)
		}
	}

	for _, v := range ids {
		fmt.Println(v)
	}

	return nil
}

func main() {
	client, err := version.New()
	if err != nil {
		log.Panic(err)
	}

	var cmd struct {
		Download DownloadCmd `cmd`
		Versions VersionsCmd `cmd`
	}

	ctx := kong.Parse(&cmd)
	err = ctx.Run(&context{client: client})

	ctx.FatalIfErrorf(err)
}
