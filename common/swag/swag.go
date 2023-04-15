package swag

import (
	"log"
	"os"

	httpSwagger "zero-template/common/swag/http"

	"github.com/swaggo/swag"
	"github.com/zeromicro/go-zero/rest"
)

type s struct {
	File string
}

func (s *s) ReadDoc() string {
	data, err := os.ReadFile(s.File)
	if err != nil {
		log.Fatal(err)
	}
	doc := string(data)
	return doc
}

func WithSwaggerHandler(file, url string) rest.RunOption {
	swag.Register(swag.Name, &s{
		File: file,
	})

	return rest.WithNotFoundHandler(httpSwagger.Handler(
		httpSwagger.Prefix(url),
		httpSwagger.InstanceName(swag.Name),
	))
}

// Deprecated
func WithDocHandler(name, url string) rest.RunOption {
	return rest.WithNotFoundHandler(httpSwagger.Handler(
		httpSwagger.Prefix(url),
		httpSwagger.InstanceName(name),
	))
}
