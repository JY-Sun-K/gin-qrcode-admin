package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"io/ioutil"
	"os"
)


func CreateQrCode(data, user string, width, height int) error{

	filename := user+".png"
	// Create the barcode
	qrCode, _ := qr.Encode(data, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, width, height)

	// create the output file
	file, _ := os.Create("./qrcodepng/"+filename)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	//fpt, err := filepath.Abs("/qrcodepng/"+filename)
	_, err := ioutil.ReadFile("./qrcodepng/"+filename)
	//imgdata, err := ioutil.ReadFile(fpt)

	if err != nil {

		return err
	}

	return nil
}