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

const (
	iconSize  = 64
	outerR    = float64(iconSize) / 2.0
	innerR    = outerR - 5.0
	centerX   = float64(iconSize) / 2.0
	centerY   = float64(iconSize) / 2.0
	arcRadius = (outerR + innerR) / 2.0
	lineWidth = outerR - innerR
)

var (
	trackColor    = color.RGBA{R: 45, G: 45, B: 45, A: 255}
	textColor     = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	chargingColor = color.RGBA{R: 46, G: 196, B: 182, A: 255}

	redColor   = color.RGBA{R: 255, G: 50, B: 50, A: 255}
	amberColor = color.RGBA{R: 255, G: 191, B: 0, A: 255}
	greenColor = color.RGBA{R: 0, G: 230, B: 0, A: 255}
)

var parsedFont *truetype.Font

func init() {
	var err error
	parsedFont, err = truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("utils: failed to parse embedded font: %v", err)
	}
}

func GenerateIcon(batteryLevel int8, isCharging bool) []byte {
	fontSize := float64(iconSize) / 1.5

	if batteryLevel < 0 {
		batteryLevel = 0
	} else if batteryLevel >= 100 {
		batteryLevel = 100
		fontSize = float64(iconSize) / 2.0
	}

	dc := gg.NewContext(iconSize, iconSize)
	dc.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 0}) // Clear background
	dc.Clear()

	// Determine active battery color
	var activeColor color.RGBA
	switch {
	case isCharging:
		activeColor = chargingColor
	case batteryLevel <= 10:
		activeColor = redColor
	case batteryLevel <= 20:
		activeColor = amberColor
	default:
		activeColor = greenColor
	}

	// Full circle track background
	dc.SetStrokeStyle(gg.NewSolidPattern(trackColor))
	dc.SetLineWidth(lineWidth)
	dc.DrawCircle(centerX, centerY, arcRadius)
	dc.Stroke()

	// Active battery level arc
	if batteryLevel > 0 {
		dc.SetStrokeStyle(gg.NewSolidPattern(activeColor))
		startAngle := gg.Radians(-90)
		endAngle := gg.Radians(-90 + 360*float64(batteryLevel)/100)
		dc.DrawArc(centerX, centerY, arcRadius, startAngle, endAngle)
		dc.Stroke()
	}

	face := truetype.NewFace(parsedFont, &truetype.Options{
		Size: fontSize,
	})
	dc.SetFontFace(face)
	dc.SetColor(textColor)
	dc.DrawStringAnchored(strconv.Itoa(int(batteryLevel)), centerX, centerY, 0.5, 0.4)

	// Pre-allocate buffer capacity to avoid slice growing overhead
	var buf bytes.Buffer
	buf.Grow(2048)

	if err := png.Encode(&buf, dc.Image()); err != nil {
		log.Printf("Error encoding PNG: %v", err)
		return nil
	}

	return buf.Bytes()
}
