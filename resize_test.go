package image

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResize(t *testing.T) {
	reader, err := os.Open(chickenImagePath)

	assert.NoError(t, err)
	defer reader.Close()

	img, err := Read(reader)
	assert.NoError(t, err)
	assert.NotEmpty(t, img)

	out, err := Resize(img, Resized(10, 10))
	assert.NoError(t, err)
	assert.NotEmpty(t, out)
}
