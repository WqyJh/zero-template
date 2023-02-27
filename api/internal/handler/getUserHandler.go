package handler

import (
	"net/http"
	"zero-template/common/httpz"

	"zero-template/api/internal/logic"
	"zero-template/api/internal/svc"
	"zero-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logc"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRequest
		if err := httpz.Parse(r, &req); err != nil {
			httpz.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		logc.Infow(r.Context(), "user handler", logc.Field("resp", resp))
		if err != nil {
			httpz.Error(w, err)
		} else {
			httpz.OkJson(w, resp)
		}
	}
}
