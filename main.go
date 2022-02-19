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
	if err := run("Babno%20Polje", "bp.png", true); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if err := run("Ljubljana", "lj.png", true); err != nil {
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
		panic(err)
	}
	err = res.DecodeField("post_id", &id)
	if err != nil {
		panic(err)
	}
	return
}

func run(location, out string, publish bool) error {
	dc := gg.NewContext(1200, 690)
	render.VremeImage(location, dc)
	if err := dc.SavePNG(out); err != nil {
		return errors.Wrap(err, "save png")
	}
	if publish {
		fbPublishPhoto(out)
	}
	return nil
}
