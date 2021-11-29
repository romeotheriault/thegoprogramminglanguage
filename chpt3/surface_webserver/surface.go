// Surface computes an SVG rendering of a 3-D surface function.
// See page 58.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defWidth, defHeight = 600, 320    // canvas size in pixels
	cells               = 100         // number of grid cells
	xyrange             = 30.0        // axis ranges (-xyrange..+xyrange)
	angle               = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		height = defHeight
	}
	width, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil {
		width = defWidth
	}
	color := r.URL.Query().Get("color")
	if color == "" {
		color = "grey"
	}
	xyscale := float64(width) / 2 / xyrange    // pixels per x or y unit
	var zscale float64 = float64(height) * 0.4 // pixels per z unit
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, height, width, xyscale, zscale)
			bx, by := corner(i, j, height, width, xyscale, zscale)
			cx, cy := corner(i, j+1, height, width, xyscale, zscale)
			dx, dy := corner(i+1, j+1, height, width, xyscale, zscale)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}

func corner(i int, j int, height int, width int, xyscale float64, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
