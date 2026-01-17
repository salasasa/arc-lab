// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"auth-rpc/pb/auth"
	"context"
	"time"

	"gateway-api/internal/svc"
	"gateway-api/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	rpcResp, err := l.svcCtx.AuthRpc.Login(l.ctx, &auth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	// 如果 RPC 返回错误（比如密码错误），直接把错误甩回给前端
	if err != nil {
		return nil, err
	}

	// 验证成功，签发 JWT
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	// 使用 RPC 返回的真实 Username 和 UserId 来签发 Token
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, rpcResp.Username)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		Role:         "admin", // TODO这里以后也可以根据 RPC 返回的权限字段动态赋值
	}, nil

}

// 辅助函数：生成 JWT 字符串
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, username string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
