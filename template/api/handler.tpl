package {{.PkgName}}

import (
    "zero-template/common/httpz"
	"net/http"

	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpz.Parse(r, &req); err != nil {
			httpz.Error(w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpz.Error(w, err)
		} else {
			{{if .HasResp}}httpz.OkJson(w, resp){{else}}httpz.Ok(w){{end}}
		}
	}
}