package squares

import (
	"image"
	"net/http"

	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// Random handler for "/squares/random"
// generates a random 6 by 6 grid image.
func Random(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)
	colors := extract.Colors(r)

	prob := extract.Probability(r, 1/float64(len(colors)))

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		squares.RandomGrid(m, colors, 6, prob)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.RandomGridSVG(w, colors, size, size, 6, prob)
	}
}
