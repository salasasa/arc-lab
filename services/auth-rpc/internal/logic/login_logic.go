package logic

import (
	"context"
	"errors"

	"auth-rpc/internal/svc"
	"auth-rpc/pb/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证登录
func (l *LoginLogic) Login(in *auth.LoginRequest) (*auth.LoginResponse, error) {
	// todo: add your logic here and delete this line
	l.Infof("RPC Login 收到请求: 用户名=%s", in.Username)

	if in.Username == "admin" && in.Password == "123456" {
		l.Info("登录成功！")
		return &auth.LoginResponse{
			UserId:   1001,
			Username: "admin",
		}, nil
	}

	l.Errorf("登录失败: 密码错误或用户不存在")
	return nil, errors.New("Invalid credentials")
}
