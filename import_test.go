package image

import (
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/stretchr/testify/assert"
)

var (
	chickenImagePath = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "chicken.jpg")
	bananaImagePath  = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "banana.png")
)

func TestImport(t *testing.T) {
	img, err := read(chickenImagePath, Width(10), Height(10))
	assert.NoError(t, err)
	assert.NotEmpty(t, img)
}
