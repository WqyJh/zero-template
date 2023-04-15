package auth

import (
	"context"
	"time"

	"zero-template/api/internal/svc"
	"zero-template/api/internal/types"
	"zero-template/common/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type JwtTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJwtTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JwtTokenLogic {
	return &JwtTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JwtTokenLogic) JwtToken(req *types.JwtTokenReq) (resp *types.JwtTokenReply, err error) {
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := jwt.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, req.UserId, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtTokenReply{
		Jwt: types.Jwt{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}
