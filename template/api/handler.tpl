package {{.PkgName}}

import (
    "chatbot/common/httpz"
	"net/http"

    "github.com/zeromicro/go-zero/core/logc"

	{{.ImportPackages}}
)

{{.HandlerComment}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpz.Parse(r, &req); err != nil {
			httpz.ErrorLog(r.Context(), w, err)
			return
		}

		logc.Infow(r.Context(), "{{.HandlerName}}", logc.Field("req", &req))

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
            httpz.Error(w, err)
			logc.Infow(r.Context(), "{{.HandlerName}} error", logc.Field("error", err))
		} else {
			{{if .HasResp}}httpz.OkJson(w, resp){{else}}httpz.Ok(w){{end}}{{if .HasResp}}
            logc.Infow(r.Context(), "{{.HandlerName}} success", logc.Field("resp", resp)){{end}}
		}
	}
}