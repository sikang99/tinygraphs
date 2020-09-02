package squares

import (
	"image"
	"net/http"

	"github.com/taironas/tinygraphs/draw/squares"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/format"
	"github.com/taironas/tinygraphs/write"
)

// BannerRandom handler for "/squares/banner/random"
// generates a random banner grid image.
func BannerRandom(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)
	xsquares := extract.XSquares(r)
	colors := extract.Colors(r)
	prob := extract.Probability(r, 1/float64(len(colors)))

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, width, height))
		squares.RandomGrid(m, colors, xsquares, prob)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.RandomGridSVG(w, colors, width, height, xsquares, prob)
	}
}

// BannerRandomGradient handler for "/squares/banner/random/gradient"
// generates a random banner grid image with gradient colors from brighter to darker color.
func BannerRandomGradient(w http.ResponseWriter, r *http.Request) {
	width := extract.Width(r)
	height := extract.Height(r)
	xsquares := extract.XSquares(r)
	colors := extract.Colors(r)

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, width, height))
		squares.RandomGradientGrid(m, colors, xsquares)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.RandomGradientGridSVG(w, colors, width, height, xsquares)
	}
}
