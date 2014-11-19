package main

import (
	"fmt"
	"io"
	"image"
	"image/png"
	"os"
	"log"
)

func main() {
	fmt.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	GenerateQRCode(file, "555-2368")
}

func GenerateQRCode(w io.Writer, code string) error {
	img := image.NewNRGBA(image.Rect(0,0, 21,21))
	return png.Encode(w, img)
}
