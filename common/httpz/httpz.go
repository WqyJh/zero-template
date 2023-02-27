package httpz

import (
	"net/http"
	"reflect"

	"zero-template/common/errorz"
	"zero-template/common/validatz"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Parse(r *http.Request, v interface{}) error {
	if err := httpx.ParsePath(r, v); err != nil {
		return err
	}

	if err := httpx.ParseForm(r, v); err != nil {
		return err
	}

	if err := httpx.ParseHeaders(r, v); err != nil {
		return err
	}

	if err := httpx.ParseJsonBody(r, v); err != nil {
		return err
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

func Ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

type responseData struct {
	errorz.CodeErrorResponse
	Data any `json:"data,omitempty"`
}

func OkJson(w http.ResponseWriter, v interface{}) {
	WriteJson(w, http.StatusOK, v)
}

func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	var res = responseData{Data: v}
	if v != nil && (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		// When v is not nil, but v.Value is nil, omitempty of Data field won't work and
		// `null` would be produced. Set the data to nil to avoid it and omit this field.
		res.Data = nil
		// Or you can use following code to convert `null` to `{}`
		// res.Data = map[string]string{}
	}
	httpx.WriteJson(w, code, res)
}
