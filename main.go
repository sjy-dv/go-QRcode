package main

import (
	"fmt"

	qrcode "github.com/yeqown/go-qrcode"
)

func main() {
	//event page url?
	qrc, err := qrcode.New("https://github.com/sjy-dv")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err := qrc.Save("./qrcode/repo-qrcode.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}