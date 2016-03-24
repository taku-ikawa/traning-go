package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math"
	"os"
)

func mandelbrot_bigFloat_main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			hypot := math.Hypot(x, y)
			// 画像の点(px、py)は複素数zを表している。
			img.Set(px, py, mandelbrot_bigFloat(*big.NewFloat(hypot)))
		}
	}

	//標準出力だとうまくいかない
	//png.Encode(os.Stdout, img) // 注意： エラーを無視
	//ファイル出力
	f, _ := os.OpenFile("data/mandelbro3.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

func mandelbrot_bigFloat(z big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	var v *big.Float
	v = big.NewFloat(0)
	for n := uint8(0); n < iterations; n++ {
		v.Add(v, v)
		i, _  := v.Int64()
		if i > 2 {
			return color.RGBA{255 - contrast*n,0,0,255}
		}
	}
	return color.Black
}
