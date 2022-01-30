package main

import (
	"fmt"
	"os"
	"path"
	"vreme/render"

	"github.com/fogleman/gg"
	fb "github.com/huandu/facebook/v2"
	"github.com/pkg/errors"
)

var access_token = os.Getenv("TOKEN")

const pageId = "107670851816100"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func fbPublishPhoto(postimage string) (id string, err error) {
	if access_token == "" {
		return
	}
	file := fb.File(postimage)
	photoPoint := path.Join(pageId, "photos")
	res, err := fb.Post(photoPoint, fb.Params{
		"access_token": access_token,
		"source":       file,
	})
	if err != nil {
		return
	}
	err = res.DecodeField("post_id", &id)
	if err != nil {
		return
	}
	return
}

func run() error {
	dc := gg.NewContext(1200, 690)
	render.VremeImage("Ljubljana", dc)
	if err := dc.SavePNG("lj.png"); err != nil {
		return errors.Wrap(err, "save png")
	}
	render.VremeImage("Babno%20Polje", dc)
	if err := dc.SavePNG("bp.png"); err != nil {
		return errors.Wrap(err, "save png")
	}

	fbPublishPhoto("bp.png")
	return nil
}
