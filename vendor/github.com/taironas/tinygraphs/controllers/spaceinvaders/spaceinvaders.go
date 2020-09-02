package spaceinvaders

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/taironas/route"
	"github.com/taironas/tinygraphs/cache"
	"github.com/taironas/tinygraphs/draw/spaceinvaders"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

// SpaceInvaders handler for /spaceinvaders/:key
func SpaceInvaders(w http.ResponseWriter, r *http.Request) {
	var err error
	var key string
	if key, err = route.Context.Get(r, "key"); err != nil {
		log.Println("Unable to get 'key' value: ", err)
		key = ""
	}

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	colors := extract.Colors(r)
	size := extract.Size(r)

	if Cache.IsCached(&w, r, key, colors, size) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	write.ImageSVG(w)
	spaceinvaders.SpaceInvaders(w, key, colors, size)
}
