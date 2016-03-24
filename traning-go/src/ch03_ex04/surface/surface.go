package surface

import (
	"fmt"
	"io"
	"math"
)

var width float64 = 600
var height float64 = 320
var xyscale float64 = width / 2 / xyrange
var zscale float64 = height * 0.4
var color string = "red"

const (
	cells   = 100         // 格子のます目の数
	xyrange = 30.0        // 軸の範囲 (-xyrange..+xyrange)
	angle   = math.Pi / 6 // x、y軸の角度 (=30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

func Surface(out io.Writer, _width float64, _height float64, _color string) {
	if _width != 0 {
		fmt.Printf("%v\n", _width)
		width = _width
	}
	if _height != 0 {
		height = _height
		fmt.Printf("%v\n", _height)
	}
	if _color != "" {
		color = _color
		fmt.Printf("%v\n", _color)
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %v; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// ます目(i,j)のかどの点(x,y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)

	// (x,y,z)を2-D SVGキャンパス(sx,xy)へ等軸的に投影。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)からの距離
	if r == 0 {
		return 0
	}
	return math.Sin(r) / r
}
