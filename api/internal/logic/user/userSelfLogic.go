package user

import (
	"context"

	"zero-template/api/internal/svc"
	"zero-template/api/internal/types"
	"zero-template/common/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelfLogic {
	return &UserSelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSelfLogic) UserSelf(req *types.UserSelfReq) (resp *types.UserSelfReply, err error) {
	uid := jwt.GetUid(l.ctx)
	l.svcCtx.UserModel.FindOne(l.ctx, uid)
	return
}
