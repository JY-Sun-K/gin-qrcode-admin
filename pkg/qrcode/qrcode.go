package qrcode

import (
	"encoding/base64"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"io/ioutil"
	"os"
	"qcode/models"
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

func PushPic(user models.User) string {
	PicName := user.Name+".png"
	fileName := "./qrcodepng/"+PicName
	PicFile, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	encodeString := base64.StdEncoding.EncodeToString(PicFile)
	return encodeString
}