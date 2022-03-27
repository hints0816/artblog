package img

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
)

func ResizeImg(url string) image.Image {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	m, _ := imaging.Decode(resp.Body)
	dstImage128 := imaging.Thumbnail(m, 400, 400, imaging.CatmullRom)
	return dstImage128
}
