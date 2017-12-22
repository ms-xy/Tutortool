package server

import (
	"bytes"
	"github.com/ms-xy/Tutortool/src/bindata"
	"io"
	"path"
)

type GoBindataAssetsTemplateLoader struct {
}

func (l *GoBindataAssetsTemplateLoader) Abs(base, name string) string {
	return path.Clean(path.Join(path.Dir(base), name))
}

func (l *GoBindataAssetsTemplateLoader) Get(path string) (io.Reader, error) {
	data, err := bindata.Asset(path)
	if err != nil {
		return nil, err
	} else {
		reader := bytes.NewReader(data)
		return reader, err
	}
}
