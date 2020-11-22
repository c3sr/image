package examples

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	raiimage "github.com/c3sr/image"
	"github.com/c3sr/image/types"
)

func Get(name string, opts ...raiimage.Option) (types.Image, error) {
  name = strings.ToLower(name)
  if filepath.Ext(name) == "" {
    for _, ext := range []string{".jpg", ".png", ".gif"} {
      if img, err := Get(name + ext , opts...); err == nil {
        return img, nil
      }
    }
  }
	bts, err := ReadFile("/" + name)
	if err != nil {
		return nil, errors.Errorf("cannot read example image %v", name)
	}
	reader := bytes.NewBuffer(bts)
	return raiimage.Read(reader, opts...)
}
