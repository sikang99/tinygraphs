package write

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

// ImageTemplate defines a basic html template to insert an image
// using data:image/jpeg;base64,mydatahere.
var ImageTemplate = `<!DOCTYPE html>
<html lang="en"><head><title>{{ .Title }}</title></head>
<body><img src="data:image/jpg;base64,{{.Image}}"></body>`

// ImageWithTemplate encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func ImageWithTemplate(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

// ImageJPEG encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func ImageJPEG(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

// ImageSVG encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func ImageSVG(w http.ResponseWriter) {
	// First set the content type in the http.ResponseWriter
	// only then create and draw the svg object.
	// the order is important because the svg object is created from the http.ResponseWriter.
	// If you create first the svg object and then you try to set the content-type to svg it will
	// not work, the content type will be text/xml and you will get the message:
	// "Resource interpreted as Image but transferred with MIME type text/xml:"
	w.Header().Set("Content-Type", "image/svg+xml")
}
