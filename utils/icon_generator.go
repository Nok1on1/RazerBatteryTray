package utils

import (
	"bytes"
	"image/color"
	"image/png"
	"log"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func GenerateIcon(batteryLevel int8) []byte {
	const size = 64
	const outerR = size / 2.0
	const innerR = outerR - 4.0
	const centerX = size / 2.0
	const centerY = size / 2.0

	if batteryLevel < 0 {
		batteryLevel = 0
	} else if batteryLevel > 100 {
		batteryLevel = 100
	}

	dc := gg.NewContext(size, size)
	dc.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 0}) // Transparent background
	dc.Clear()

	// Determine active color based on battery level
	var activeColor color.RGBA
	if batteryLevel <= 10 {
		activeColor = color.RGBA{R: 255, G: 50, B: 50, A: 255} // Red
	} else if batteryLevel <= 20 {
		activeColor = color.RGBA{R: 255, G: 191, B: 0, A: 255} // Amber/Yellow
	} else {
		activeColor = color.RGBA{R: 0, G: 230, B: 0, A: 255} // Razer Green
	}

	trackColor := color.RGBA{R: 45, G: 45, B: 45, A: 255}
	textColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	// full circle track
	dc.SetStrokeStyle(gg.NewSolidPattern(trackColor))
	dc.SetLineWidth(outerR - innerR)
	dc.DrawCircle(centerX, centerY, (outerR+innerR)/2)
	dc.Stroke()

	// active battery level arc
	if batteryLevel > 0 {
		dc.SetStrokeStyle(gg.NewSolidPattern(activeColor))
		startAngle := gg.Radians(-90) // Start from top
		endAngle := gg.Radians(-90 + 360*float64(batteryLevel)/100)
		dc.DrawArc(centerX, centerY, (outerR+innerR)/2, startAngle, endAngle)
		dc.Stroke()
	}

	// Load font and draw text
	parsedFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Printf("Error parsing font: %v", err)
		return nil
	}

	face := truetype.NewFace(parsedFont, &truetype.Options{
		Size: size / 1.5,
		DPI:  72,
	})
	dc.SetFontFace(face)
	dc.SetColor(textColor)
	dc.DrawStringAnchored(strconv.Itoa(int(batteryLevel)), centerX, centerY, 0.5, 0.4)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		log.Printf("Error encoding PNG: %v", err)
		return nil
	}
	return buf.Bytes()
}
