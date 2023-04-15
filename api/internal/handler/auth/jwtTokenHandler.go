package auth

import (
	"net/http"
	"zero-template/common/httpz"

	"zero-template/api/internal/logic/auth"
	"zero-template/api/internal/svc"
	"zero-template/api/internal/types"
)

// JwtToken godoc
// @Summary      获取 Jwt Token
// @Description  获取 Jwt Token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req   body      types.JwtTokenReq  true  "请求体"
// @Success      200    {object}   types.DataResponse{data=types.JwtTokenReply}
// @Router       /api/v1/auth/token [post]
func JwtTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JwtTokenReq
		if err := httpz.Parse(r, &req); err != nil {
			httpz.Error(w, err)
			return
		}

		l := auth.NewJwtTokenLogic(r.Context(), svcCtx)
		resp, err := l.JwtToken(&req)
		if err != nil {
			httpz.Error(w, err)
		} else {
			httpz.OkJson(w, resp)
		}
	}
}
