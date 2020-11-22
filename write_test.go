package image

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/c3sr/image/types"
	"github.com/stretchr/testify/assert"
)

var (
	laneControlImagePath = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "lane_control.jpg")
)

func TestWrite(t *testing.T) {
	reader, err := os.Open(laneControlImagePath)

	assert.NoError(t, err)
	defer reader.Close()

	img0, err := Read(reader)
	assert.NoError(t, err)
	assert.NotEmpty(t, img0)

	img, ok := img0.(*types.RGBImage)
	assert.True(t, ok)

	toPng("/tmp/image_test.png", img.Pix, img.Bounds())
}

func toPng(filePath string, imgByte []byte, bounds image.Rectangle) {

	img := types.NewRGBImage(bounds)
	copy(img.Pix, imgByte)

	out, _ := os.Create(filePath)
	defer out.Close()

	err := png.Encode(out, img.ToRGBAImage())
	if err != nil {
		log.Println(err)
	}
}
