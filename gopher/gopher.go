package gopher

import (
	"bytes"
	"embed"
	"image"
	"image/draw"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	xdraw "golang.org/x/image/draw"
)

const size = 32

//go:embed images/*
var images embed.FS

var (
	Front *ebiten.Image
	Side  *ebiten.Image
	Pink  *ebiten.Image
)

func init() {
	front, _ := loadImage("images/gopher.png")
	Front = ebiten.NewImageFromImage(resizeByHeight(front))
	side, _ := loadImage("images/gopher-side.png")
	Side = ebiten.NewImageFromImage(resizeByHeight(side))
	pink, _ := loadImage("images/gopher-pink.png")
	Pink = ebiten.NewImageFromImage(resizeByHeight(pink))
}

func loadImage(path string) (image.Image, error) {
	b, err := images.ReadFile(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func resizeByHeight(img image.Image) image.Image {
	imgDst := image.NewRGBA(image.Rect(0, 0, int(float64(img.Bounds().Dx())*size/float64(img.Bounds().Dy())), size)) // heightを基準にリサイズ
	xdraw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return imgDst.SubImage(imgDst.Rect)
}
