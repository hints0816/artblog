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
	dstImage128 := imaging.Resize(m, 200, 200, imaging.Lanczos)
	imaging.Save(dstImage128, "d:/2.png")
	return dstImage128
}
