package httpz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"zero-template/common/errorz"
	"zero-template/common/validatz"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Parse(r *http.Request, v interface{}) error {
	if err := httpx.ParsePath(r, v); err != nil {
		return errorz.NewCodeError(http.StatusBadRequest, err.Error())
	}

	if err := httpx.ParseForm(r, v); err != nil {
		return errorz.NewCodeError(http.StatusBadRequest, err.Error())
	}

	if err := httpx.ParseHeaders(r, v); err != nil {
		return errorz.NewCodeError(http.StatusBadRequest, err.Error())
	}

	if err := httpx.ParseJsonBody(r, v); err != nil {
		return errorz.NewCodeError(http.StatusBadRequest, err.Error())
	}

	if msg := validatz.Validate(v, r.Header.Get("Accept-Language")); msg != "" {
		return errorz.NewCodeError(http.StatusBadRequest, msg)
	}

	return nil
}

func parseError(err error) (int, interface{}) {
	switch e := err.(type) {
	case *errorz.CodeError:
		return http.StatusOK, e.Data()
	default:
		return http.StatusInternalServerError, nil
	}
}

func Error(w http.ResponseWriter, err error) {
	code, body := parseError(err)
	if body == nil {
		w.WriteHeader(code)
		return
	}
	e, ok := body.(error)
	if ok {
		http.Error(w, e.Error(), code)
	} else {
		httpx.WriteJson(w, code, body)
	}
}

func ErrorLog(ctx context.Context, w http.ResponseWriter, err error) {
	logx.WithContext(ctx).WithCallerSkip(1).Errorv(err)
	Error(w, err)
}

func Ok(w http.ResponseWriter) {
	OkJson(w, nil)
}

type responseData struct {
	errorz.CodeErrorResponse
	Data any `json:"data,omitempty"`
}

func OkJson(w http.ResponseWriter, v interface{}) {
	WriteJson(w, http.StatusOK, v)
}

// OkJsonCtx writes v into w with 200 OK.
func OkJsonCtx(ctx context.Context, w http.ResponseWriter, v interface{}) {
	WriteJsonCtx(ctx, w, http.StatusOK, v)
}

// WriteJson writes v as json string into w with code.
func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	if err := doWriteJson(w, code, v); err != nil {
		logx.Error(err)
	}
}

// WriteJsonCtx writes v as json string into w with code.
func WriteJsonCtx(ctx context.Context, w http.ResponseWriter, code int, v interface{}) {
	if err := doWriteJson(w, code, v); err != nil {
		logx.WithContext(ctx).Error(err)
	}
}

func doWriteJson(w http.ResponseWriter, code int, v interface{}) error {
	var res = responseData{Data: v}
	res.Code = code

	if v != nil && (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		// When v is not nil, but v.Value is nil, omitempty of Data field won't work and
		// `null` would be produced. Set the data to nil to avoid it and omit this field.
		res.Data = nil
		// Or you can use following code to convert `null` to `{}`
		// res.Data = map[string]string{}
	}

	bs, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("marshal json failed, error: %w", err)
	}

	w.Header().Set(httpx.ContentType, httpx.JsonContentType)
	w.WriteHeader(code)

	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(bs) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}

	return nil
}
