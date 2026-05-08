package winres

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"testing"
)

var (
	benchmarkBytesSink []byte
	benchmarkRelocSink []int
	benchmarkImageSink image.Image
)

func BenchmarkResourceSetBytes(b *testing.B) {
	rs := ResourceSet{}
	for i := 1; i <= 24; i++ {
		data := bytes.Repeat([]byte{byte(i)}, i*37)
		_ = rs.Set(RT_RCDATA, ID(i), uint16(i), data)
		_ = rs.Set(Name(fmt.Sprintf("CUSTOM_%02d", i)), Name(fmt.Sprintf("Payload_%02d", i)), uint16(0x400+i), data)
	}

	b.ReportAllocs()
	var data []byte
	var reloc []int
	for b.Loop() {
		data, reloc = rs.bytes()
	}
	benchmarkBytesSink = data
	benchmarkRelocSink = reloc
}

func BenchmarkImageInSquareNRGBA(b *testing.B) {
	img := image.NewNRGBA(image.Rect(0, 0, 192, 96))
	for y := range img.Bounds().Dy() {
		for x := range img.Bounds().Dx() {
			img.SetNRGBA(x, y, color.NRGBA{
				R: uint8(x),
				G: uint8(y),
				B: uint8(x + y),
				A: 0xff,
			})
		}
	}

	b.ReportAllocs()
	var square image.Image
	for b.Loop() {
		square = imageInSquareNRGBA(img, true)
	}
	benchmarkImageSink = square
}
