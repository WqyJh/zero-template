package user

import (
	"net/http"
	"zero-template/common/httpz"

	"zero-template/api/internal/logic/user"
	"zero-template/api/internal/svc"
	"zero-template/api/internal/types"
)

// UserSelf godoc
// @Summary      查询用户信息
// @Description  查询用户信息
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        client   header      string  true  "当前设备类型: android/ios"
// @Param        cversion   header      int64  true  "客户端版本号：整数"
// @Param        mid   header      string  true  "设备ID"
// @Success      200    {object}   types.DataResponse{data=types.UserSelfReply}
// @Router       /api/v1/user/self [get]
func UserSelfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSelfReq
		if err := httpz.Parse(r, &req); err != nil {
			httpz.Error(w, err)
			return
		}

		l := user.NewUserSelfLogic(r.Context(), svcCtx)
		resp, err := l.UserSelf(&req)
		if err != nil {
			httpz.Error(w, err)
		} else {
			httpz.OkJson(w, resp)
		}
	}
}
