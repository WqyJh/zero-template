package restz

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func AddHealthCheck(server *rest.Server, path string) {
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   path,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	})
}
