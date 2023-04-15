package jwt

import (
	"context"
	"encoding/json"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

const UidKey = "uid"

func GenToken(iat int64, secretKey string, userId int64, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[UidKey] = userId

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func GetUid(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(UidKey).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorw("GetUid err", logx.Field("error", err))
		}
	}
	return uid
}
