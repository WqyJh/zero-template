package restz

import (
	"net/http"
	"zero-template/common/errorz"
	"zero-template/common/httpz"
)

// Response 200 OK {"code": 401, "msg": "error msg"}
// instead of just response 401 Unauthorized
func UnauthorizedCallback(w http.ResponseWriter, r *http.Request, err error) {
	httpz.Error(w, errorz.NewCodeError(401, err.Error()))
}
