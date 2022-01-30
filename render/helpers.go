package render

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"time"
	"vreme/arso"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

func getPNG(icon string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/render/pngs/%s.png", dir, icon)
}

func cc(dc *gg.Context) {
	margin := 20.0
	x := margin
	y := margin

	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 20); err != nil {
		panic(err)
	}
	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(100),
	}
	dc.SetColor(mutedColor)
	s := "vir vreme.si"
	marginX := 50.0
	marginY := 25.0
	textWidth, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)

}

func kraj(tekst string, dc *gg.Context) {
	margin := 20.0
	x := margin + 30
	y := margin + 130

	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 80); err != nil {
		panic(err)
	}
	dc.SetColor(color.Black)
	dc.DrawString(tekst, x+1, y+1)
	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(180),
	}
	dc.SetColor(mutedColor)
	dc.DrawString(tekst, x, y)
}
func sedaj(icon, t string, dc *gg.Context) {
	text := fmt.Sprintf("%s°C", t)
	w, _ := dc.MeasureString(text)
	x := float64(dc.Width()) - float64(w) - float64(180)
	y := float64(190)

	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 120); err != nil {
		panic(err)
	}

	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(180),
	}
	dc.SetColor(mutedColor)
	dc.DrawString(text, x, y)

	//Icon
	f, _ := gg.LoadImage(getPNG(icon))
	f = imaging.Fill(f, 120, 120, imaging.Center, imaging.Lanczos)
	dc.DrawImage(f, int(x)-140, int(y)-105)
}

func vzhodZahod(vz, zh string, dc *gg.Context) {
	tekst := fmt.Sprintf("Sonce vzhaja %s, zahaja %s", vz, zh)
	margin := 20.0
	x := margin + 38
	y := margin + 180

	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 24); err != nil {
		panic(err)
	}
	dc.SetColor(color.Black)
	dc.DrawString(tekst, x+1, y+1)
	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(220),
	}
	dc.SetColor(mutedColor)
	dc.DrawString(tekst, x, y)
}

func veter(dir string, dc *gg.Context) {
	tekst := fmt.Sprintf("Veter: %s", dir)
	margin := 20.0
	x := margin + 38
	y := margin + 210

	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 24); err != nil {
		panic(err)
	}
	dc.SetColor(color.Black)
	dc.DrawString(tekst, x+1, y+1)
	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(220),
	}
	dc.SetColor(mutedColor)
	dc.DrawString(tekst, x, y)
}

func vlaga(rr string, dc *gg.Context) {
	tekst := fmt.Sprintf("Vlaga: %s", rr)
	margin := 20.0
	x := margin + 38
	y := margin + 240

	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 24); err != nil {
		panic(err)
	}
	dc.SetColor(color.Black)
	dc.DrawString(tekst, x+1, y+1)
	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(220),
	}
	dc.SetColor(mutedColor)
	dc.DrawString(tekst, x, y)
}

func napoved(datum string, low, high string, img string, split int, dc *gg.Context) {
	full := dc.Height()
	top := 25
	//head
	fontPath := filepath.Join("OpenSans", "OpenSans-SemiBold.ttf")
	if err := dc.LoadFontFace(fontPath, 40); err != nil {
		panic(err)
	}
	textColor := color.White
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(150),
	}
	dc.SetColor(mutedColor)
	dc.DrawString(datum, float64(60+split), float64((full-380)+top))

	//Icon
	t, _ := gg.LoadImage(img)
	t = imaging.Fill(t, 90, 90, imaging.Center, imaging.Lanczos)
	dc.DrawImage(t, 65+split, full-350+top)

	//high
	if err := dc.LoadFontFace(fontPath, 30); err != nil {
		panic(err)
	}
	dc.SetColor(mutedColor)
	dc.DrawString(fixT(fmt.Sprintf("%s°C", low)), float64(75+split), float64(full-195+top))
	dc.DrawString(fixT(fmt.Sprintf("%s°C", high)), float64(75+split), float64(full-135+top))

}

func fixT(in string) string {
	if len(in) <= 4 {
		return fmt.Sprintf(" %s", in)
	}
	return in
}

func ozadje(dc *gg.Context) {
	backgroundImageFilename := "jutro.jpeg"
	if time.Now().UTC().Hour() > 10 {
		backgroundImageFilename = "dan.jpeg"
	}
	if time.Now().UTC().Hour() > 18 {
		backgroundImageFilename = "vecer.jpeg"
	}
	backgroundImage, err := gg.LoadImage(backgroundImageFilename)
	if err != nil {
		panic(err)
	}
	backgroundImage = imaging.Fill(backgroundImage, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)
	dc.DrawImage(backgroundImage, 0, 0)
}

func overlay(dc *gg.Context) {
	margin := 30.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 180})

	if time.Now().UTC().Hour() > 18 {
		dc.SetColor(color.RGBA{0, 0, 0, 150})
	}
	dc.DrawRoundedRectangle(x, y, w, h, 30)
	dc.Fill()
}

func VremeImage(loc string, dc *gg.Context) {

	ozadje(dc)
	overlay(dc)
	data := arso.GetArso(loc)
	today := data.Observation.Features[0].Properties.Days[0]
	now := today.Timeline[0]
	forecast := data.Forecast24H.Features[0].Properties.Days[1:8]

	kraj(data.Observation.Features[0].Properties.Title, dc)
	sedaj(now.CloudsIconWwsynIcon, now.T, dc)
	vzhodZahod(today.Sunrise.Format("15:04"), today.Sunset.Format("15:04"), dc)
	veter(fmt.Sprintf("%s, hitrost %skm/h", now.FfShortText, now.FfVal), dc)
	vlaga(fmt.Sprintf("%s %s%%", now.RhShortText, now.Rh), dc)

	sep := 160

	for index, day := range forecast {
		lsep := 0
		if index >= 1 {
			lsep = sep * index
		}
		date := strings.Split(day.Date, "-")
		napoved(
			fmt.Sprintf("%s.%s", date[2], date[1]),
			day.Timeline[0].Tnsyn,
			day.Timeline[0].Txsyn,
			getPNG(day.Timeline[0].CloudsIconWwsynIcon),
			lsep,
			dc,
		)
	}

	cc(dc)
}
