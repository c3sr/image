package examples

import (
	"bytes"

	"github.com/pkg/errors"
	raiimage "github.com/rai-project/image"
	"github.com/rai-project/image/types"
)

func Get(name string, opts ...raiimage.Option) (types.Image, error) {
	bts, err := ReadFile("/" + name)
	if err != nil {
		return nil, errors.Errorf("cannot read example image %v", name)
	}
	reader := bytes.NewBuffer(bts)
	return raiimage.Read(reader, opts...)
}
