package qrcode

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func ResizeImgw() {
	qrCode, _ := qr.Encode("http://www.topgoer.com", qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create("qr2.png")
	defer file.Close()

	png.Encode(file, qrCode)
}
