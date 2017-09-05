package image

// import (
// 	"path/filepath"
// 	"testing"

// 	context "golang.org/x/net/context"

// 	sourcepath "github.com/GeertJohan/go-sourcepath"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	chickenImagePath = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "chicken.jpg")
// 	bananaImagePath  = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "banana.png")
// )

// func TestImport1(t *testing.T) {
// 	img, err := read(context.Background(), chickenImagePath, Width(10), Height(10))
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, img)
// }

// func TestImport2(t *testing.T) {
// 	img, err := read(context.Background(), bananaImagePath, Width(10), Height(10))
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, img)
// }
