package util

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

const UTF8_BOTH = "\xE2\x96\x88"
const UTF8_TOPHALF = "\xE2\x96\x80"
const UTF8_BOTTOMHALF = "\xE2\x96\x84"

func PrintQRCode(content string) {
	qr, err := qrcode.New(content, qrcode.Low)
	if err != nil {
		log.Fatal(err)
	}

	bitmap := qr.Bitmap()
	for i := 0; i < len(bitmap); i += 2 {
		for j := 0; j < len(bitmap[i]); j++ {
			if i+1 < len(bitmap) {
				if bitmap[i][j] && bitmap[i+1][j] {
					fmt.Print(UTF8_BOTH)
				} else if bitmap[i][j] {
					fmt.Print(UTF8_TOPHALF)
				} else if bitmap[i+1][j] {
					fmt.Print(UTF8_BOTTOMHALF)
				} else {
					fmt.Print(" ")
				}
			} else {
				if bitmap[i][j] {
					fmt.Print(UTF8_TOPHALF)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
