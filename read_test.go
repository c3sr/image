package image

import (
	"os"
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/stretchr/testify/assert"
)

var (
	chickenImagePath = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "chicken.jpg")
	bananaImagePath  = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "banana.png")
)

func TestImport1(t *testing.T) {
	reader, err := os.Open(chickenImagePath)

	assert.NoError(t, err)
	defer reader.Close()

	img, err := Read(reader)
	assert.NoError(t, err)
	assert.NotEmpty(t, img)
}

func TestImport2(t *testing.T) {
	reader, err := os.Open(bananaImagePath)

	assert.NoError(t, err)
	defer reader.Close()

	img, err := Read(reader)
	assert.NoError(t, err)
	assert.NotEmpty(t, img)
}
